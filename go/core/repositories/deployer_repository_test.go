package repositories

import (
    "time";
    "testing";

    "github.com/marekgalovic/photon/go/core/storage";

    "github.com/stretchr/testify/suite";
    "github.com/samuel/go-zookeeper/zk";
)

type DeployerRepositoryTest struct {
    suite.Suite
    zk *storage.Zookeeper
    repository *deployerRepository
}

func TestDeployerRepository(t *testing.T) {
    suite.Run(t, new(DeployerRepositoryTest))
}

func (test *DeployerRepositoryTest) SetupSuite() {
    test.zk = storage.NewTestZookeeper()
    test.repository = NewDeployerRepository(test.zk)

    test.seedZookeeper()
}

func (test *DeployerRepositoryTest) TearDownSuite() {
    test.zk.Close()
}

func (test *DeployerRepositoryTest) seedZookeeper() {
    _, err := test.zk.Create("/deployer/pmml/models/model_uid/versions/version_uid", "/path/version_uid", int32(1), zk.WorldACL(zk.PermAll))
    test.Require().Nil(err)

    _, err = test.zk.Create("/deployer/pmml/models/model_uid/versions/version_uid2", "/path/version_uid2", int32(1), zk.WorldACL(zk.PermAll))
    test.Require().Nil(err)
}

func (test *DeployerRepositoryTest) TestListVersionsW() {
    versions, event, err := test.repository.ListVersionsW("pmml", "model_uid")
    test.Require().Nil(err)

    test.NotNil(event)
    test.Equal(map[string]string{"version_uid": "/path/version_uid", "version_uid2": "/path/version_uid2"}, versions)
}

func (test *DeployerRepositoryTest) TestListVersionsWFiresAfterVersionsAreChanged() {
    sig := make(chan bool, 1)
    go func() {
        select {
        case <- sig:
            err := test.zk.Delete("/deployer/pmml/models/model_uid/versions/version_uid", -1)
            test.Require().Nil(err)
        case <- time.After(1 * time.Second):
            return
        }
    }()
    _, event, err := test.repository.ListVersionsW("pmml", "model_uid")
    test.Require().Nil(err)

    sig <- true

    select {
    case <- event:
        versions, _, err := test.repository.ListVersionsW("pmml", "model_uid")
        test.Require().Nil(err)
        test.Equal(map[string]string{"version_uid2": "/path/version_uid2"}, versions)
    case <- time.After(1 * time.Second):
        test.FailNow("Timeout while waiting for watch event.")
    }
}

func (test *DeployerRepositoryTest) TestModelPathExistsWFirestAfterModelPathIsCreated() {
    sig := make(chan bool, 1)
    go func() {
        select {
        case <- sig:
            _, err := test.zk.Create("/deployer/pmml/models/model_uid2/versions", nil, int32(1), zk.WorldACL(zk.PermAll))
            test.Require().Nil(err)
        case <- time.After(1 * time.Second):
            return
        }
    }()
    _, event, err := test.repository.ModelPathExistsW("pmml", "model_uid2")
    test.Require().Nil(err)

    sig <- true

    select {
    case <- event:
        return
    case <- time.After(1 * time.Second):
        test.FailNow("Timeout while waiting for watch event.") 
    }
}

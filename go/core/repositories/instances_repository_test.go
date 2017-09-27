package repositories

import (
    "time";
    "testing";

    "github.com/marekgalovic/photon/go/core/storage";

    "github.com/stretchr/testify/suite";
    "github.com/samuel/go-zookeeper/zk";
)

type InstancesRepositoryTest struct {
    suite.Suite
    zk *storage.Zookeeper
    repository *instancesRepository
}

func TestInstancesRepository(t *testing.T) {
    suite.Run(t, new(InstancesRepositoryTest))
}

func (test *InstancesRepositoryTest) SetupTest() {
    test.zk = storage.NewTestZookeeper()
    test.repository = NewInstancesRepository(test.zk)

    test.seedZookeeper()
}

func (test *InstancesRepositoryTest) TearDownTest() {
    test.zk.Close()
}

func (test *InstancesRepositoryTest) seedZookeeper() {
    _, err := test.zk.Create("/instances/model_version_uid_a/instance_uid_a", map[string]interface{}{"address": "127.0.0.1", "port": 5022}, int32(1), zk.WorldACL(zk.PermAll))
    test.Require().Nil(err)

    _, err = test.zk.Create("/instances/model_version_uid_a/instance_uid_b", map[string]interface{}{"address": "0.0.0.0", "port": 5023}, int32(1), zk.WorldACL(zk.PermAll))
    test.Require().Nil(err)
}

func (test *InstancesRepositoryTest) TestList() {
    instances, err := test.repository.List("model_version_uid_a")

    test.Require().Nil(err)
    test.Require().Equal(2, len(instances))
    test.Equal(&Instance{Address: "127.0.0.1", Port: 5022}, instances["instance_uid_a"])
    test.Equal(&Instance{Address: "0.0.0.0", Port: 5023}, instances["instance_uid_b"])
}

func (test *InstancesRepositoryTest) TestRegister() {
    uid, err := test.repository.Register("model_version_uid_a", &Instance{Address: "127.0.0.1", Port: 5006})
    test.Require().Nil(err)

    children, err := test.zk.Children("/instances/model_version_uid_a")
    test.Require().Nil(err)

    for _, child := range children {
        if child == uid {
            return
        }
    }
    test.FailNow("Registered children not found.")
}

func (test *InstancesRepositoryTest) TestUnregister() {
    err := test.repository.Unregister("model_version_uid_a", "instance_uid_a")
    test.Require().Nil(err)

    children, err := test.zk.Children("/instances/model_version_uid_a")

    test.Require().Nil(err)
    test.Equal(1, len(children))
}

func (test *InstancesRepositoryTest) TestWatch() {
    createSig := make(chan bool, 1)
    go func() {
        select {
        case <- createSig:
            _, err := test.zk.Create("/instances/model_version_uid_a/instance_uid_c", map[string]interface{}{"address": "0.0.0.0", "port": 5024}, int32(1), zk.WorldACL(zk.PermAll))
            test.Require().Nil(err)
        case <- time.After(1 * time.Second):
            return
        }
    }()

    event, err := test.repository.Watch("model_version_uid_a")
    test.Require().Nil(err)

    createSig <- true

    select {
    case <- event:
        instances, err := test.repository.List("model_version_uid_a")
        test.Require().Nil(err)
        test.Equal(3, len(instances))
    case <- time.After(1 * time.Second):
        test.FailNow("Timeout while waiting for watch event.")
    }
}

func (test *InstancesRepositoryTest) TestExists() {
    exists, err := test.repository.Exists("model_version_uid_a")

    test.Nil(err)
    test.True(exists)
}

package balancer

import (
    "time";
    "testing";

    "github.com/marekgalovic/photon/server/storage";
    "github.com/marekgalovic/photon/server/storage/repositories";

    "github.com/stretchr/testify/suite";
    "github.com/samuel/go-zookeeper/zk";
)

type WatcherTest struct {
    suite.Suite
    zk *storage.Zookeeper
    repository *repositories.InstancesRepository
    watcher *Watcher
    availableInstances map[string]*repositories.Instance
}

func TestWatcher(t *testing.T) {
    suite.Run(t, new(WatcherTest))
}

func (test *WatcherTest) SetupSuite() {
    test.zk = storage.NewTestZookeeper()
    test.repository = repositories.NewInstancesRepository(test.zk)

    test.seedZookeeper()
}

func (test *WatcherTest) TearDownSuite() {
    test.zk.Close()
}

func (test *WatcherTest) SetupTest() {
    test.watcher = NewWatcher(test.repository, "model_version_uid")
}

func (test *WatcherTest) TearDownTest() {
    test.watcher.Close()
}

func (test *WatcherTest) seedZookeeper() {
    test.availableInstances = map[string]*repositories.Instance{
        "instance_uid_a": &repositories.Instance{Uid: "instance_uid_a", Address: "127.0.0.1", Port: 5022},
        "instance_uid_b": &repositories.Instance{Uid: "instance_uid_b", Address: "127.0.0.1", Port: 5023},
        "instance_uid_c": &repositories.Instance{Uid: "instance_uid_c", Address: "127.0.0.1", Port: 5024},
    }

    _, err := test.zk.Create("/instances/model_version_uid/instance_uid_a", map[string]interface{}{"address": "127.0.0.1", "port": 5022}, int32(1), zk.WorldACL(zk.PermAll))
    test.Require().Nil(err)

    _, err = test.zk.Create("/instances/model_version_uid/instance_uid_b", map[string]interface{}{"address": "127.0.0.1", "port": 5023}, int32(1), zk.WorldACL(zk.PermAll))
    test.Require().Nil(err)

    _, err = test.zk.Create("/instances/model_version_uid/instance_uid_c", map[string]interface{}{"address": "127.0.0.1", "port": 5024}, int32(1), zk.WorldACL(zk.PermAll))
    test.Require().Nil(err)
}

func (test *WatcherTest) TestNext() {
    updates, err := test.watcher.Next()

    test.Require().Nil(err)
    test.Equal(3, len(updates))
}

func (test *WatcherTest) TestNextReturnsOnChange() {
    createSig := make(chan bool, 1)
    go func() {
        select {
        case <- createSig:
            _, err := test.zk.Create("/instances/model_version_uid/instance_uid_d", map[string]interface{}{"address": "127.0.0.1", "port": 5025}, int32(1), zk.WorldACL(zk.PermAll))
            test.Require().Nil(err)
        case <- time.After(1 * time.Second):
            return
        }
    }()

    _, err := test.watcher.Next()
    test.Require().Nil(err)

    createSig <- true

    updates, err := test.watcher.Next()
    test.Require().Nil(err)
    test.Equal(1, len(updates))
}

package server

import (
    "time";
    "testing";

    "github.com/marekgalovic/photon/server/storage";
    "github.com/marekgalovic/photon/server/storage/repositories";

    "github.com/stretchr/testify/suite";
    "github.com/samuel/go-zookeeper/zk";
)

type InstanceResolverTest struct {
    suite.Suite
    zk *storage.Zookeeper
    repository *repositories.InstancesRepository
    resolver *InstanceResolver
    availableInstances map[string]*repositories.Instance
}

func TestInstanceResolver(t *testing.T) {
    suite.Run(t, new(InstanceResolverTest))
}

func (test *InstanceResolverTest) SetupSuite() {
    test.zk = storage.NewTestZookeeper()
    test.repository = repositories.NewInstancesRepository(test.zk)

    test.seedZookeeper()
}

func (test *InstanceResolverTest) TearDownSuite() {
    test.zk.Close()
}

func (test *InstanceResolverTest) SetupTest() {
    test.resolver = NewInstanceResolver(test.repository)
}

func (test *InstanceResolverTest) seedZookeeper() {
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

func (test *InstanceResolverTest) TestGet() {
    instance, err := test.resolver.Get("model_version_uid")
    test.Require().Nil(err)

    expectedValue, exists := test.availableInstances[instance.Uid]
    test.Equal(true, exists)
    test.Equal(expectedValue, instance)
}

func (test *InstanceResolverTest) TestGetCachesFetchedInstances() {
    _, err := test.resolver.Get("model_version_uid")
    test.Require().Nil(err)

    cacheData, exists := test.resolver.instancesCache.Get("model_version_uid")
    test.Require().Equal(true, exists)
    test.Equal(3, len(cacheData.([]*repositories.Instance)))
}

func (test *InstanceResolverTest) TestWatchUpdatesInstancesCacheOnChange() {
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

    _, err := test.resolver.Get("model_version_uid")
    test.Require().Nil(err)

    createSig <- true
    time.Sleep(100 * time.Millisecond)

    cacheData, exists := test.resolver.instancesCache.Get("model_version_uid")
    test.Require().Equal(true, exists)
    test.Equal(4, len(cacheData.([]*repositories.Instance)))  
}

func (test *InstanceResolverTest) TestGetLoadBalancing() {
    instanceA, err := test.resolver.Get("model_version_uid")
    test.Require().Nil(err)

    instanceB, err := test.resolver.Get("model_version_uid")
    test.Require().Nil(err)

    instanceC, err := test.resolver.Get("model_version_uid")
    test.Require().Nil(err)

    instanceD, err := test.resolver.Get("model_version_uid")
    test.Require().Nil(err)

    test.NotEqual(instanceA, instanceB)
    test.NotEqual(instanceA, instanceC)
    test.NotEqual(instanceB, instanceC)
    test.Equal(instanceA, instanceD)
}

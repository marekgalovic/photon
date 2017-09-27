package balancer

import (
    "testing";

    "github.com/marekgalovic/photon/go/core/storage";
    "github.com/marekgalovic/photon/go/core/repositories";

    "github.com/stretchr/testify/suite";
    "github.com/samuel/go-zookeeper/zk";
)

type ResolverTest struct {
    suite.Suite
    zk *storage.Zookeeper
    repository repositories.InstancesRepository
    resolver *Resolver
}

func TestResolver(t *testing.T) {
    suite.Run(t, new(ResolverTest))
}

func (test *ResolverTest) SetupSuite() {
    test.zk = storage.NewTestZookeeper()
    test.resolver = NewResolver(repositories.NewInstancesRepository(test.zk))
}

func (test *ResolverTest) TearDownSuite() {
    test.zk.Close()
}

func (test *ResolverTest) seedZookeeper() {
    _, err := test.zk.Create("/instances/model_version_uid/instance_uid_a", map[string]interface{}{"address": "127.0.0.1", "port": 5022}, int32(1), zk.WorldACL(zk.PermAll))
    test.Require().Nil(err)
}

func (test *ResolverTest) TestResolve() {
    watcher, err := test.resolver.Resolve("model_version_uid")

    test.Require().Nil(err)
    watcher.Close()
}

func (test *ResolverTest) TestResolveNonExistingVersion() {
    watcher, err := test.resolver.Resolve("model_version_foo")
    
    test.NotNil(err)
    test.Nil(watcher)
}

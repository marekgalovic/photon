package providers

import (
    "github.com/marekgalovic/photon/server/storage"
)

type ZookeeperProvider struct {
    zk *storage.Zookeeper
}

func NewZookeeperProvider(zk *storage.Zookeeper) *ZookeeperProvider {
    return &ZookeeperProvider{zk: zk}
}

func (z *ZookeeperProvider) Get(modelVersionUid string) (*Instance, error) {
    return &Instance{Address: "127.0.0.1", Port: 7001}, nil
}

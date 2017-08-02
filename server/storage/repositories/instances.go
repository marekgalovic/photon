package repositories

import (
    "github.com/marekgalovic/photon/server/storage"
)

type Instance struct {

}

type InstancesRepository struct {
    zk *storage.Zookeeper
}

func NewInstancesRepository(zk *storage.Zookeeper) *InstancesRepository {
    return &InstancesRepository{zk: zk}
}

func (r *InstancesRepository) List(modelUid string) ([]*Instance, error) {
    return nil, nil
}

func (r *InstancesRepository) Get(modelUid, versionUid string) (*Instance, error) {
    return nil, nil
}

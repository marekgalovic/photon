package repositories

import (
    "path/filepath";

    "github.com/marekgalovic/photon/server/storage";

    "github.com/samuel/go-zookeeper/zk";
)

type Instance struct {
    Uid string
    Address string
    Port int
}

type InstancesRepository struct {
    zk *storage.Zookeeper
}

func NewInstancesRepository(zk *storage.Zookeeper) *InstancesRepository {
    return &InstancesRepository{zk: zk}
}

func (r *InstancesRepository) List(versionUid string) ([]*Instance, error) {
    children, err := r.zk.ChildrenData(r.instancesPath(versionUid))
    if err != nil {
        return nil, err
    }

    return r.scanInstances(children)
}

func (r *InstancesRepository) Watch(versionUid string) (<-chan zk.Event, error) {
    _, event, err := r.zk.ChildrenW(r.instancesPath(versionUid))

    return event, err
}

func (r *InstancesRepository) instancesPath(versionUid string) string {
    return filepath.Join("instances", versionUid)
}

func (r *InstancesRepository) scanInstances(children map[string]*storage.ZNode) ([]*Instance, error) {
    instances := make([]*Instance, 0, len(children))

    for name, znode := range children {
        instance := &Instance{Uid: name}

        if err := znode.Scan(&instance); err != nil {
            return nil, err
        }
        
        instances = append(instances, instance)
    }

    return instances, nil
}

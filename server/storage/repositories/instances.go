package repositories

import (
    "path/filepath";

    "github.com/marekgalovic/photon/server/storage";
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

func (r *InstancesRepository) ListW(versionUid string) ([]*Instance, error) {
    _, event, err := r.zk.ChildrenW(r.instancesPath(versionUid))
    if err != nil {
        return nil, err
    }

    <- event
    return r.List(versionUid)
}

func (r *InstancesRepository) instancesPath(versionUid string) string {
    return filepath.Join("instances", versionUid)
}

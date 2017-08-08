package repositories

import (
    "fmt";
    "path/filepath";

    "github.com/marekgalovic/photon/go/core/storage";

    "github.com/samuel/go-zookeeper/zk";
    "github.com/satori/go.uuid";
)

type Instance struct {
    Uid string 
    Address string
    Port int
}

func (i *Instance) FullAddress() string {
    return fmt.Sprintf("%s:%d", i.Address, i.Port)
}

type InstancesRepository struct {
    zk *storage.Zookeeper
}

func NewInstancesRepository(zk *storage.Zookeeper) *InstancesRepository {
    return &InstancesRepository{zk: zk}
}

func (r *InstancesRepository) List(versionUid string) (map[string]*Instance, error) {
    children, err := r.zk.ChildrenData(r.instancesPath(versionUid))
    if err != nil {
        return nil, err
    }

    return r.scanInstances(children)
}

func (r *InstancesRepository) Register(versionUid, address string, port int) (string, error) {
    name := filepath.Join(r.instancesPath(versionUid), fmt.Sprintf("%s", uuid.NewV4()))

    fullPath, err := r.zk.CreateEphemeralSequential(name, &Instance{Address: address, Port: port}, zk.WorldACL(zk.PermRead))
    if err != nil {
        return "", err
    }

    return filepath.Base(fullPath), nil
}

func (r *InstancesRepository) Unregister(versionUid, uid string) error {
    path := filepath.Join(r.instancesPath(versionUid), uid)

    return r.zk.Delete(path, -1)
}

func (r *InstancesRepository) Watch(versionUid string) (<-chan zk.Event, error) {
    _, event, err := r.zk.ChildrenW(r.instancesPath(versionUid))

    return event, err
}

func (r *InstancesRepository) scanInstances(children map[string]*storage.ZNode) (map[string]*Instance, error) {
    instances := make(map[string]*Instance)

    for uid, znode := range children {
        instance := &Instance{}
        if err := znode.Scan(&instance); err != nil {
            return nil, err
        }
        instance.Uid = uid
        instances[uid] = instance
    }

    return instances, nil
}

func (r *InstancesRepository) instancesPath(versionUid string) string {
    return filepath.Join("instances", versionUid)
}

package repositories

import (
    "fmt";
    "path/filepath";

    "github.com/marekgalovic/photon/go/core/storage";

    "github.com/samuel/go-zookeeper/zk";
    "gopkg.in/go-playground/validator.v9"
)

type Instance struct {
    Address string `validate:"required"`
    Port int `validate:"required,gt=1000"`
}

func (i *Instance) FullAddress() string {
    return fmt.Sprintf("%s:%d", i.Address, i.Port)
}

type InstancesRepository interface {
    List(string) (map[string]*Instance, error)
    Register(string, *Instance) (string, error)
    Unregister(string, string) error
    Watch(string) (<-chan zk.Event, error)
    Exists(string) (bool, error)
}

type instancesRepository struct {
    zk *storage.Zookeeper
    validate *validator.Validate
}

func NewInstancesRepository(zk *storage.Zookeeper) *instancesRepository {
    return &instancesRepository{
        zk: zk,
        validate: validator.New(),
    }
}

func (r *instancesRepository) List(versionUid string) (map[string]*Instance, error) {
    children, err := r.zk.ChildrenData(r.instancesPath(versionUid))
    if err != nil {
        return nil, err
    }

    return r.scanInstances(children)
}

func (r *instancesRepository) Register(versionUid string, instance *Instance) (string, error) {
    if err := r.validate.Struct(instance); err != nil {
        return "", err
    }
    name := filepath.Join(r.instancesPath(versionUid), "runner")

    fullPath, err := r.zk.CreateEphemeralSequential(name, instance, zk.WorldACL(zk.PermRead))
    if err != nil {
        return "", err
    }

    return filepath.Base(fullPath), nil
}

func (r *instancesRepository) Unregister(versionUid, uid string) error {
    path := filepath.Join(r.instancesPath(versionUid), uid)

    return r.zk.Delete(path, -1)
}

func (r *instancesRepository) Watch(versionUid string) (<-chan zk.Event, error) {
    _, event, err := r.zk.ChildrenW(r.instancesPath(versionUid))

    return event, err
}

func (r *instancesRepository) Exists(versionUid string) (bool, error) {
    exists, err := r.zk.Exists(r.instancesPath(versionUid))

    return exists, err
}

func (r *instancesRepository) scanInstances(children map[string]*storage.ZNode) (map[string]*Instance, error) {
    instances := make(map[string]*Instance)

    for uid, znode := range children {
        instance := &Instance{}
        if err := znode.Scan(&instance); err != nil {
            return nil, err
        }
        instances[uid] = instance
    }

    return instances, nil
}

func (r *instancesRepository) instancesPath(versionUid string) string {
    return filepath.Join("instances", versionUid)
}

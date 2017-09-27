package balancer

import (
    "fmt";
    
    "github.com/marekgalovic/photon/go/core/repositories";

    "google.golang.org/grpc/naming";
)

type Resolver struct {
    instancesRepository repositories.InstancesRepository
}

func NewResolver(instancesRepository repositories.InstancesRepository) *Resolver {
    return &Resolver{
        instancesRepository: instancesRepository,
    }
}

func (r *Resolver) Resolve(versionUid string) (naming.Watcher, error) {
    exists, err := r.instancesRepository.Exists(versionUid)
    if err != nil {
        return nil, err
    }
    if !exists {
        return nil, fmt.Errorf("Cannot resolve instances for version '%s'.", versionUid)
    }
    return NewWatcher(r.instancesRepository, versionUid), nil
}

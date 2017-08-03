package balancer

import (
    "github.com/marekgalovic/photon/server/storage/repositories";

    "google.golang.org/grpc/naming";
)

type Resolver struct {
    instancesRepository *repositories.InstancesRepository
}

func NewResolver(instancesRepository *repositories.InstancesRepository) *Resolver {
    return &Resolver{
        instancesRepository: instancesRepository,
    }
}

func (r *Resolver) Resolve(modelVersionUid string) (naming.Watcher, error) {
    return NewWatcher(r.instancesRepository, modelVersionUid), nil
}

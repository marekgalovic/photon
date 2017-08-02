package server

import (
    "github.com/marekgalovic/photon/server/storage/repositories"
)

type InstanceResolver struct {
    instancesRepository *repositories.InstancesRepository
}

func NewInstanceResolver(instancesRepository *repositories.InstancesRepository) *InstanceResolver {
    return &InstanceResolver{
        instancesRepository: instancesRepository,
    }
}

func (r *InstanceResolver) Get(modelUid, modelVersionUid string) (*repositories.Instance, error) {
    return nil, nil
}

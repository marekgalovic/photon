package server

import (
    "time";
    "math/rand";

    "github.com/marekgalovic/photon/server/storage/repositories";

    "github.com/patrickmn/go-cache";
    log "github.com/Sirupsen/logrus"
)

type InstanceResolver struct {
    instancesRepository *repositories.InstancesRepository
    instancesCache *cache.Cache
    roundRobinIndex *cache.Cache
    instanceWatches *cache.Cache
}

func NewInstanceResolver(instancesRepository *repositories.InstancesRepository) *InstanceResolver {
    return &InstanceResolver{
        instancesRepository: instancesRepository,
        instancesCache: cache.New(10 * time.Minute, 10 * time.Minute),
        roundRobinIndex: cache.New(10 * time.Minute, 10 * time.Minute),
        instanceWatches: cache.New(10 * time.Minute, 10 * time.Minute),
    }
}

func (r *InstanceResolver) Get(modelVersionUid string) (*repositories.Instance, error) {
    if cached, exists := r.instancesCache.Get(modelVersionUid); exists {
        return r.roundRobin(modelVersionUid, cached.([]*repositories.Instance)), nil
    }

    instances, err := r.instancesRepository.List(modelVersionUid)
    if err != nil {
        return nil, err
    }
    r.instancesCache.Set(modelVersionUid, instances, cache.DefaultExpiration)

    if _, exists := r.instanceWatches.Get(modelVersionUid); !exists {
        go r.watchInstances(modelVersionUid)
    }

    return r.roundRobin(modelVersionUid, instances), nil
}

func (r *InstanceResolver) watchInstances(modelVersionUid string) {
    defer r.instanceWatches.Delete(modelVersionUid)
    r.instanceWatches.Set(modelVersionUid, nil, cache.NoExpiration)

    watchLogger := log.WithFields(log.Fields{"model_version_uid": modelVersionUid})
    watchLogger.Debug("Setting instances watch.")
    for {
        event, err := r.instancesRepository.Watch(modelVersionUid)
        if err != nil {
            watchLogger.Errorf("Set instances watch failed: %v", err)
            return
        }

        <- event

        instances, err := r.instancesRepository.List(modelVersionUid)
        if err != nil {
            watchLogger.Errorf("Fetch updated instances failed: %v", err)
            return
        }

        watchLogger.Debug("Instances updated.")
        r.instancesCache.Set(modelVersionUid, instances, cache.DefaultExpiration)
    }
}

func (r *InstanceResolver) roundRobin(modelVersionUid string, instances []*repositories.Instance) *repositories.Instance {
    var index int

    if cached, exists := r.roundRobinIndex.Get(modelVersionUid); exists {
        index = cached.(int)
    } else {
        index = rand.Intn(len(instances))
    }

    if index > len(instances) - 1 {
        index = 0
    }

    r.roundRobinIndex.Set(modelVersionUid, index + 1, cache.DefaultExpiration)

    return instances[index]
}

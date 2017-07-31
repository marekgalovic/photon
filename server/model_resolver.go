 package server

import (
    "fmt";
    "time";
    // "math/rand";

    "github.com/marekgalovic/photon/server/metrics";
    "github.com/marekgalovic/photon/server/storage/repositories";

    "github.com/patrickmn/go-cache";
)

type ModelResolver struct {
    modelsRepository *repositories.ModelsRepository
    modelsCache *cache.Cache
}

type modelsCacheEntry struct {
    model *repositories.Model
    version *repositories.ModelVersion
}   

func NewModelResolver(modelsRepository *repositories.ModelsRepository) *ModelResolver {
    return &ModelResolver{
        modelsRepository: modelsRepository,
        modelsCache: cache.New(30 * time.Second, 1 * time.Minute),
    }
}

func (m *ModelResolver) GetModel(uid string) (*repositories.Model, *repositories.ModelVersion, error) {
    defer metrics.Runtime("model_resolver.get_model.runtime", []string{fmt.Sprintf("model_uid:%s", uid)})

    if cached, exists := m.modelsCache.Get(uid); exists {
        entry := cached.(*modelsCacheEntry)
        return entry.model, entry.version, nil
    }

    model, err := m.modelsRepository.Find(uid)
    if err != nil {
        return nil, nil, err
    }

    version, err := m.modelsRepository.PrimaryVersion(model.Uid)
    if err != nil {
        return nil, nil, err
    }

    m.modelsCache.Set(uid, &modelsCacheEntry{model: model, version: version}, cache.DefaultExpiration)
    return model, version, nil
}

func (m *ModelResolver) GetShadowModels(uid string) ([]*repositories.ModelVersion, error) {
    return nil, nil
}

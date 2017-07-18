 package server

import (
    "fmt";
    "time";
    // "math/rand";

    "github.com/marekgalovic/photon/server/metrics";
    "github.com/marekgalovic/photon/server/storage";
)

type ModelResolver struct {
    modelsRepository *storage.ModelsRepository
    modelsCache map[string]*modelsCacheEntry
    modelsCacheTimeout time.Duration
}

type modelsCacheEntry struct {
    cachedAt time.Time
    model *storage.Model
    version *storage.ModelVersion
}   

func NewModelResolver(modelsRepository *storage.ModelsRepository) *ModelResolver {
    return &ModelResolver{
        modelsRepository: modelsRepository,
        modelsCache: make(map[string]*modelsCacheEntry, 0),
        modelsCacheTimeout: 10 * time.Second,
    }
}

func (m *ModelResolver) GetModel(uid string) (*storage.Model, *storage.ModelVersion, error) {
    defer metrics.Runtime("model_resolver.get_model.runtime", []string{fmt.Sprintf("model_uid:%s", uid)})

    if cached, exists := m.modelsCache[uid]; exists && time.Since(cached.cachedAt) < m.modelsCacheTimeout {
        return cached.model, cached.version, nil
    }

    model, err := m.modelsRepository.Find(uid)
    if err != nil {
        return nil, nil, err
    }

    version, err := m.modelsRepository.PrimaryVersion(model.Uid)
    if err != nil {
        return nil, nil, err
    }

    m.modelsCache[uid] = &modelsCacheEntry{cachedAt: time.Now(), model: model, version: version} 
    return model, version, nil
}

func (m *ModelResolver) GetShadowModels(uid string) ([]*storage.Model, error) {
    return nil, nil
}

package evaluator

import (
    "fmt";
    "time";

    "github.com/marekgalovic/photon/go/core/metrics";
    "github.com/marekgalovic/photon/go/core/repositories";

    "github.com/patrickmn/go-cache";
)

type ModelResolver interface {
    GetModel(string) (*repositories.Model, *repositories.ModelVersion, error)
}

type modelResolver struct {
    modelsRepository repositories.ModelsRepository
    modelsCache *cache.Cache
}

type modelsCacheEntry struct {
    model *repositories.Model
    version *repositories.ModelVersion
}   

func NewModelResolver(modelsRepository repositories.ModelsRepository) *modelResolver {
    return &modelResolver{
        modelsRepository: modelsRepository,
        modelsCache: cache.New(30 * time.Second, 1 * time.Minute),
    }
}

func (m *modelResolver) GetModel(name string) (*repositories.Model, *repositories.ModelVersion, error) {
    defer metrics.Runtime("model_resolver.runtime", []string{fmt.Sprintf("model_name:%s", name), "method:get_model"})

    if cached, exists := m.modelsCache.Get(name); exists {
        entry := cached.(*modelsCacheEntry)
        return entry.model, entry.version, nil
    }

    model, err := m.modelsRepository.FindByName(name)
    if err != nil {
        return nil, nil, err
    }

    version, err := m.modelsRepository.PrimaryVersion(model.Id)
    if err != nil {
        return nil, nil, err
    }

    m.modelsCache.Set(name, &modelsCacheEntry{model: model, version: version}, cache.DefaultExpiration)
    return model, version, nil
}

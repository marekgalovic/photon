 package server

import (
    // "math/rand";

    // "github.com/marekgalovic/serving/server/metrics";
    "github.com/marekgalovic/serving/server/storage";
)

type ModelResolver struct {
    modelsRepository *storage.ModelsRepository
}

func NewModelResolver(modelsRepository *storage.ModelsRepository) *ModelResolver {
    return &ModelResolver{
        modelsRepository: modelsRepository,
    }
}

func (m *ModelResolver) GetModel(uid string) (*storage.Model, *storage.ModelVersion, error) {
    model, err := m.modelsRepository.Find(uid)
    if err != nil {
        return nil, nil, err
    }

    version, err := m.modelsRepository.PrimaryVersion(model.Uid)
    if err != nil {
        return nil, nil, err
    }

    return model, version, nil
}

func (m *ModelResolver) GetShadowModels(uid string) ([]*storage.Model, error) {
    return nil, nil
}

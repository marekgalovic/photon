package server

import (
    // "math/rand";

    "github.com/marekgalovic/serving/server/storage"
)

type ModelManager struct {
    modelsRepository *storage.ModelsRepository
}

func NewModelManager(modelsRepository *storage.ModelsRepository) *ModelManager {
    return &ModelManager{modelsRepository: modelsRepository}
}

func (m *ModelManager) GetModel(uid string) (*storage.Model, error) {
    model, err := m.modelsRepository.Find(uid)
    if err != nil {
        return nil, err
    }

    return model, nil
}

func (m *ModelManager) GetShadowModels(uid string) ([]*storage.Model, error) {
    return nil, nil
}

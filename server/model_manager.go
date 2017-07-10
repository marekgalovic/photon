package server

import (
    "github.com/marekgalovic/serving/server/storage"
)

type ModelManager struct {
    db *storage.ModelsRepository
}

func NewModelManager(db *storage.ModelsRepository) *ModelManager {
    return &ModelManager{db: db}
}

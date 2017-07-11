package storage

import (
    "time";
)

type Model struct {
    Uid string
    Name string
    Owner string
    CreatedAt time.Time
    UpdatedAt time.Time
    PrimaryVersion *ModelVersion
}

type ModelVersion struct {
    Uid string
    Name string
    RequestFeatures []string
    StoredFeatures []string
}

type ModelsRepository struct {

}

func NewModelsRepository() *ModelsRepository {
    return &ModelsRepository{}
}

func (r *ModelsRepository) Find(uid string) (*Model, error) {
    return &Model{
        Uid: "x000",
        Name: "fraud",
        Owner: "risk.algorithms@shopify.com",
        PrimaryVersion: &ModelVersion{
            Uid: "x001",
            Name: "Random Forest 11 Features",
            RequestFeatures: []string{"x1", "x2"},
            StoredFeatures: []string{"x3", "x4"},
        },
    }, nil
}

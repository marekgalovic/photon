package server

import (
    "github.com/marekgalovic/serving/server/storage";
)

type FeaturesResolver struct {
    db *storage.FeaturesRepository
}

func NewFeaturesResolver(db *storage.FeaturesRepository) *FeaturesResolver {
    return &FeaturesResolver{db: db}
}

func (r *FeaturesResolver) Resolve(modelVersion *storage.ModelVersion, requestFeatures map[string]interface{}) (map[string]interface{}, error) {
    return requestFeatures, nil
}

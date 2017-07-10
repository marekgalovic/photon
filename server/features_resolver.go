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

func (r *FeaturesResolver) Resolve(features []string) (map[string]interface{}, error) {
    return nil, nil
}

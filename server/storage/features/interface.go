package features

import (
    "github.com/marekgalovic/photon/server/storage/repositories";
)

type FeaturesStore interface {
    Get(*repositories.FeatureSet, map[string]interface{}) (map[string]interface{}, error)
    Insert(*repositories.FeatureSet, *repositories.FeatureSetSchema, map[string]interface{}) error
    CreateFeatureSet(*repositories.FeatureSet) error
    DeleteFeatureSet(string) error
}

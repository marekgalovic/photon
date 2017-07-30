package features

import (
    "github.com/marekgalovic/photon/server/storage";
)

type FeaturesStore struct {
    db *storage.Cassandra
}

func NewFeaturesStore(db *storage.Cassandra) *FeaturesStore {
    return &FeaturesStore{
        db: db,
    }
}

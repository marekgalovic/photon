package features

import (
    "github.com/marekgalovic/photon/server/storage";
)

type CassandraFeaturesStore struct {
    db *storage.Cassandra
}

func NewCassandraFeaturesStore(db *storage.Cassandra) *CassandraFeaturesStore {
    return &CassandraFeaturesStore{
        db: db,
    }
}

func (s *CassandraFeaturesStore) Get(name string, fields []string) (map[string]interface{}, error) {
    return nil, nil
}

func (s *CassandraFeaturesStore) CreateFeaturesSet(name string) error {
    return nil
}

func (s *CassandraFeaturesStore) DeleteFeaturesSet(name string) error {
    return nil
}

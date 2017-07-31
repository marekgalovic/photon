package features

import (
    "fmt";
    "strings";
    "encoding/json";

    "github.com/marekgalovic/photon/server/storage";
    "github.com/marekgalovic/photon/server/storage/repositories";
)

type CassandraFeaturesStore struct {
    db *storage.Cassandra
}

func NewCassandraFeaturesStore(db *storage.Cassandra) *CassandraFeaturesStore {
    return &CassandraFeaturesStore{
        db: db,
    }
}

func (s *CassandraFeaturesStore) Get(featureSet *repositories.FeatureSet, params map[string]interface{}) (map[string]interface{}, error) {
    sql := fmt.Sprintf(
        `SELECT data FROM %s WHERE %s LIMIT 1`,
        s.normalizeName(featureSet.Uid), strings.Join(s.selectConditions(featureSet.Keys, params), " AND "),
    )

    var marshaledData []byte
    if err := s.db.Query(sql).Scan(&marshaledData); err != nil {
        return nil, err
    }

    var data map[string]interface{}
    if err := json.Unmarshal(marshaledData, &data); err != nil {
        return nil, err
    }

    return data, nil
}

func (s *CassandraFeaturesStore) Insert(featureSet *repositories.FeatureSet, schema *repositories.FeatureSetSchema, data map[string]interface{}) error {
    sql := fmt.Sprintf(
        `INSERT INTO %s (schema_uid,%s,data) VALUES (%s)`,
        s.normalizeName(featureSet.Uid), strings.Join(featureSet.Keys, ","), strings.TrimSuffix(strings.Repeat("?,", len(featureSet.Keys)+2), ","),
    )

    values, err := s.insertValues(schema.Uid, featureSet.Keys, data)
    if err != nil {
        return nil
    }

    return s.db.Query(sql, values...).Exec()
}

func (s *CassandraFeaturesStore) CreateFeatureSet(uid string, keys []string) error {
    sql := fmt.Sprintf(
        `CREATE TABLE %s (schema_uid UUID, %s, data TEXT, PRIMARY KEY (%s))`,
        s.normalizeName(uid), strings.Join(s.keysSchema(keys), ","), strings.Join(keys, ","),
    )

    return s.db.Query(sql).Exec()
}

func (s *CassandraFeaturesStore) DeleteFeatureSet(uid string) error {
    sql := fmt.Sprintf("DROP TABLE %s", s.normalizeName(uid))

    return s.db.Query(sql).Exec()
}

func (s *CassandraFeaturesStore) selectConditions(keys []string, params map[string]interface{}) []string {
    conditions := make([]string, len(keys))
    for _, key := range keys {
        conditions = append(conditions, fmt.Sprintf("%s = '%v'", key, params[key]))
    }
    return conditions
}

func (s *CassandraFeaturesStore) insertValues(schema_uid string, keys []string, data map[string]interface{}) ([]interface{}, error) {
    values := []interface{}{schema_uid}

    for _, key := range keys {
        value, exists := data[key]
        if !exists {
            return nil, fmt.Errorf("Missing key '%s'", key)
        }
        values = append(values, fmt.Sprintf("%v", value))
    }

    marshaledData, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }

    values = append(values, marshaledData)

    return values, nil
}

func (s *CassandraFeaturesStore) keysSchema(keys []string) []string {
    schema := make([]string, 0, len(keys))
    for _, key := range keys {
        schema = append(schema, fmt.Sprintf("%s VARCHAR", key))
    }
    return schema
}

func (s *CassandraFeaturesStore) normalizeName(name string) string {
    return strings.Replace(name, "-", "_", -1) 
}

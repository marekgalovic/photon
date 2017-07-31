package features

import (
    "fmt";
    "strings";
    "encoding/json";

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

func (s *CassandraFeaturesStore) Get(uid string, keys map[string]interface{}) (map[string]interface{}, error) {
    sql := fmt.Sprintf(
        `SELECT data FROM %s WHERE %s LIMIT 1`,
        s.normalizeName(uid), strings.Join(s.selectConditions(keys), " AND "),
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

func (s *CassandraFeaturesStore) Insert(uid, schema_uid string, keys []string, data map[string]interface{}) error {
    sql := fmt.Sprintf(
        `INSERT INTO %s (schema_uid,%s,data) VALUES (%s)`,
        s.normalizeName(uid), strings.Join(keys, ","), strings.TrimSuffix(strings.Repeat("?,", len(keys)+2), ","),
    )

    values, err := s.insertValues(schema_uid, keys, data)
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

func (s *CassandraFeaturesStore) selectConditions(keys map[string]interface{}) []string {
    conditions := make([]string, 0)
    for key, value := range keys {
        conditions = append(conditions, fmt.Sprintf("%s = '%v'", key, value))
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

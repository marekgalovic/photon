package features

import (
    "fmt";
    "strings";
    "encoding/json";

    "github.com/marekgalovic/photon/go/core/storage";
    "github.com/marekgalovic/photon/go/core/metrics";
)

type cassandraFeaturesStore struct {
    db *storage.Cassandra
}

func NewCassandraFeaturesStore(db *storage.Cassandra) *cassandraFeaturesStore {
    return &cassandraFeaturesStore{
        db: db,
    }
}

func (s *cassandraFeaturesStore) Get(id int64, keys []string, params map[string]interface{}) (map[string]interface{}, error) {
    defer metrics.Runtime("features_store.runtime", []string{"type:cassandra", "method:get"})()

    sql := fmt.Sprintf(
        `SELECT data FROM %s WHERE %s LIMIT 1`,
        s.TableName(id), strings.Join(s.selectConditions(keys, params), " AND "),
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

func (s *cassandraFeaturesStore) Insert(id int64, keys []string, data map[string]interface{}) error {
    defer metrics.Runtime("features_store.runtime", []string{"type:cassandra", "method:insert"})()

    sql := fmt.Sprintf(
        `INSERT INTO %s (%s,data) VALUES (%s)`,
        s.TableName(id), strings.Join(keys, ","), strings.TrimSuffix(strings.Repeat("?,", len(keys)+1), ","),
    )

    values, err := s.insertValues(keys, data)
    if err != nil {
        return nil
    }

    return s.db.Query(sql, values...).Exec()
}

func (s *cassandraFeaturesStore) CreateFeatureSet(id int64, keys []string) error {
    defer metrics.Runtime("features_store.runtime", []string{"type:cassandra", "method:create_feature_set"})()

    sql := fmt.Sprintf(
        `CREATE TABLE %s (%s, data TEXT, PRIMARY KEY (%s))`,
        s.TableName(id), strings.Join(s.keysSchema(keys), ","), strings.Join(keys, ","),
    )

    return s.db.Query(sql).Exec()
}

func (s *cassandraFeaturesStore) DeleteFeatureSet(id int64) error {
    defer metrics.Runtime("features_store.runtime", []string{"type:cassandra", "method:delete_feature_set"})()

    sql := fmt.Sprintf("DROP TABLE %s", s.TableName(id))

    return s.db.Query(sql).Exec()
}

func (s *cassandraFeaturesStore) TableName(featureSetId int64) string {
    return fmt.Sprintf("feature_set_%d", featureSetId)     
}

func (s *cassandraFeaturesStore) selectConditions(keys []string, params map[string]interface{}) []string {
    conditions := make([]string, len(keys))
    for i, key := range keys {
        conditions[i] = fmt.Sprintf("%s = '%v'", key, params[key])
    }
    return conditions
}

func (s *cassandraFeaturesStore) insertValues(keys []string, data map[string]interface{}) ([]interface{}, error) {
    values := make([]interface{}, 0)

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

func (s *cassandraFeaturesStore) keysSchema(keys []string) []string {
    schema := make([]string, 0, len(keys))
    for _, key := range keys {
        schema = append(schema, fmt.Sprintf("%s VARCHAR", key))
    }
    return schema
}

package storage

import (
    "fmt";
    "time";
    "strings";

    "github.com/marekgalovic/serving/server/metrics";

    "github.com/satori/go.uuid";
)

type FeatureSet struct {
    Uid string
    Name string
    Keys []string
    CreatedAt time.Time
    UpdatedAt time.Time
}

type FeatureSetSchema struct {
    Uid string
    Fields []*FeatureSetSchemaField
    CreatedAt time.Time
}

type FeatureSetSchemaField struct {
    Name string
    ValueType string
    Nullable bool
}

type FeaturesRepository struct {
    db *Mysql
}

func NewFeaturesRepository(db *Mysql) *FeaturesRepository {
    return &FeaturesRepository{db: db}
}

func (r *FeaturesRepository) List() ([]*FeatureSet, error) {
    defer metrics.Runtime("queries.runtime", []string{"repository:features", "query:list"})()

    rows, err := r.db.Query(`SELECT uid, name, lookup_keys, created_at, updated_at FROM feature_sets ORDER BY updated_at DESC`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    featureSets := make([]*FeatureSet, 0)
    for rows.Next() {
        featureSet, err := r.scanFeatureSet(rows)
        if err != nil {
            return nil, err
        }
        featureSets = append(featureSets, featureSet)
    }

    return featureSets, nil
}

func (r *FeaturesRepository) Find(uid string) (*FeatureSet, error) {
    defer metrics.Runtime("queries.runtime", []string{"repository:features", "query:find"})()

    row, err := r.db.QueryRowPrepared(`SELECT uid, name, lookup_keys, created_at, updated_at FROM feature_sets WHERE uid = ?`, uid)
    if err != nil {
        return nil, err
    }

    return r.scanFeatureSet(row)
}

func (r *FeaturesRepository) Create(name string, keys []string) (*FeatureSet, error) {
    if len(keys) < 1 {
        return nil, fmt.Errorf("Cannot create feature set with no keys")
    }
    defer metrics.Runtime("queries.runtime", []string{"repository:features", "query:create"})()

    uid := fmt.Sprintf("%s", uuid.NewV4())

    _, err := r.db.ExecPrepared(`INSERT INTO feature_sets (uid, name, lookup_keys) VALUES (?,?,?)`, uid, name, strings.Join(keys, ","))
    if err != nil {
        return nil, err
    }

    return r.Find(uid)
}

func (r *FeaturesRepository) Delete(uid string) error {
    defer metrics.Runtime("queries.runtime", []string{"repository:features", "query:delete"})()

    _, err := r.db.ExecPrepared(`DELETE FROM feature_sets WHERE uid = ?`, uid)
    
    return err
}

func (r *FeaturesRepository) scanFeatureSet(rows Scannable) (*FeatureSet, error) {
    featureSet := &FeatureSet{}
    var keys string

    if err := rows.Scan(&featureSet.Uid, &featureSet.Name, &keys, &featureSet.CreatedAt, &featureSet.UpdatedAt); err != nil {
        return nil, err
    }
    featureSet.Keys = strings.Split(keys, ",")

    return featureSet, nil
}

func (r *FeaturesRepository) ListSchemas(setUid string) ([]*FeatureSetSchema, error) {
    defer metrics.Runtime("queries.runtime", []string{"repository:features", "query:list_schemas"})()

    rows, err := r.db.QueryPrepared(`SELECT uid, created_at FROM feature_set_schemas WHERE feature_set_uid = ? ORDER BY created_at DESC`, setUid)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    schemas := make([]*FeatureSetSchema, 0)
    for rows.Next() {
        schema, err := r.scanFeatureSetSchema(rows)
        if err != nil {
            return nil, err
        }
        schemas = append(schemas, schema)
    }

    return schemas, nil
}

func (r *FeaturesRepository) FindSchema(uid string) (*FeatureSetSchema, error) {
    defer metrics.Runtime("queries.runtime", []string{"repository:features", "query:find_schema"})()

    row, err := r.db.QueryRowPrepared(`SELECT uid, created_at FROM feature_set_schemas WHERE uid = ?`, uid)
    if err != nil {
        return nil, err
    }
    schema, err := r.scanFeatureSetSchema(row)
    if err != nil {
        return nil, err
    }
    schema.Fields, err = r.schemaFields(schema.Uid)
    if err != nil {
        return nil, err
    }

    return schema, nil
}

func (r *FeaturesRepository) LatestSchema(featureSetUid string) (*FeatureSetSchema, error) {
    defer metrics.Runtime("queries.runtime", []string{"repository:features", "query:latest_schema"})()

    row, err := r.db.QueryRowPrepared(`SELECT uid, created_at FROM feature_set_schemas WHERE feature_set_uid = ? ORDER BY created_at DESC LIMIT 1`, featureSetUid)
    if err != nil {
        return nil, err
    }
    schema, err := r.scanFeatureSetSchema(row)
    if err != nil {
        return nil, err
    }
    schema.Fields, err = r.schemaFields(schema.Uid)
    if err != nil {
        return nil, err
    }

    return schema, nil
}

func (r *FeaturesRepository) CreateSchema(setUid string, fields []*FeatureSetSchemaField) (*FeatureSetSchema, error) {
    if len(fields) < 1 {
        return nil, fmt.Errorf("Cannot create feature set schema with no fields")
    }
    defer metrics.Runtime("queries.runtime", []string{"repository:features", "query:create_schema"})()

    uid := fmt.Sprintf("%s", uuid.NewV4())

    tx, err := r.db.Begin()
    if err != nil {
        return nil, err
    }

    createSetSchemaStmt, err := tx.Prepare(`INSERT INTO feature_set_schemas (uid, feature_set_uid) VALUES (?,?)`)
    if err != nil {
        return nil, err
    }
    defer createSetSchemaStmt.Close()

    if _, err = createSetSchemaStmt.Exec(uid, setUid); err != nil {
        return nil, err
    }

    createSchemaFieldsStmt, err := tx.Prepare(fmt.Sprintf(`INSERT INTO feature_set_schema_fields (feature_set_schema_uid, name, value_type, nullable) VALUES %s`, strings.TrimSuffix(strings.Repeat("(?,?,?,?),", len(fields)), ",")))
    if err != nil {
        return nil, err
    }
    defer createSchemaFieldsStmt.Close()

    if _, err = createSchemaFieldsStmt.Exec(r.schemaFieldsValues(uid, fields)...); err != nil {
        return nil, err
    }

    if err = tx.Commit(); err != nil {
        return nil, err
    }

    return r.FindSchema(uid)
}

func (r *FeaturesRepository) schemaFieldsValues(schemaUid string, fields []*FeatureSetSchemaField) []interface{} {
    values := make([]interface{}, 0, len(fields)*3)
    for _, field := range fields {
        values = append(values, schemaUid, field.Name, field.ValueType, field.Nullable)
    }
    return values
}

func (r *FeaturesRepository) DeleteSchema(uid string) error {
    defer metrics.Runtime("queries.runtime", []string{"repository:features", "query:delete_schema"})()

    _, err := r.db.ExecPrepared(`DELETE FROM feature_set_schemas WHERE uid = ?`, uid)
    
    return err
}

func (r *FeaturesRepository) schemaFields(schemaUid string) ([]*FeatureSetSchemaField, error) {
    rows, err := r.db.Query(`SELECT name, value_type, nullable FROM feature_set_schema_fields WHERE feature_set_schema_uid = ?`, schemaUid)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    fields := make([]*FeatureSetSchemaField, 0)
    for rows.Next() {
        field := &FeatureSetSchemaField{}
        if err := rows.Scan(&field.Name, &field.ValueType, &field.Nullable); err != nil {
            return nil, err
        }
        fields = append(fields, field)
    }

    return fields, nil
}

func (r *FeaturesRepository) scanFeatureSetSchema(rows Scannable) (*FeatureSetSchema, error) {
    schema := &FeatureSetSchema{}
    if err := rows.Scan(&schema.Uid, &schema.CreatedAt); err != nil {
        return nil, err
    }
    return schema, nil
}

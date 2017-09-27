package repositories

import (
    "fmt";
    "time";
    "strings";
    "database/sql";

    "github.com/marekgalovic/photon/go/core/storage";
    "github.com/marekgalovic/photon/go/core/metrics";

    "gopkg.in/go-playground/validator.v9"
)

type FeatureSet struct {
    Id int64
    Name string `validate:"required"`
    Keys []string `validate:"required,gt=0"`
    Fields []*FeatureSetField `validate:"required,gt=0,dive"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

type FeatureSetField struct {
    Id int64
    FeatureSetId int64
    Name string `validate:"required"`
    ValueType string `validate:"required"`
    Nullable bool
}

type FeaturesRepository interface {
    List() ([]*FeatureSet, error)
    Find(int64) (*FeatureSet, error)
    Create(*FeatureSet) (int64, error)
    Delete(int64) error
}

type featuresRepository struct {
    db *storage.Mysql
    validate *validator.Validate
}

func NewFeaturesRepository(db *storage.Mysql) *featuresRepository {
    return &featuresRepository{
        db: db,
        validate: validator.New(),
    }
}

func (r *featuresRepository) List() ([]*FeatureSet, error) {
    defer metrics.Runtime("queries.runtime", []string{"repository:features", "query:list"})()

    rows, err := r.db.Query(`SELECT id, name, lookup_keys, created_at, updated_at FROM feature_sets ORDER BY updated_at DESC`)
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

func (r *featuresRepository) Find(id int64) (*FeatureSet, error) {
    defer metrics.Runtime("queries.runtime", []string{"repository:features", "query:find"})()

    row, err := r.db.QueryRowPrepared(`SELECT id, name, lookup_keys, created_at, updated_at FROM feature_sets WHERE id = ?`, id)
    if err != nil {
        return nil, err
    }

    featureSet, err := r.scanFeatureSet(row)
    if err != nil {
        return nil, err
    }

    if featureSet.Fields, err = r.queryFeatureSetFields(featureSet.Id); err != nil {
        return nil, err
    }

    return featureSet, nil
}

func (r *featuresRepository) Create(featureSet *FeatureSet) (int64, error) {
    if err := r.validate.Struct(featureSet); err != nil {
        return 0, err
    }
    defer metrics.Runtime("queries.runtime", []string{"repository:features", "query:create"})()

    tx, err := r.db.Begin()
    if err != nil {
        return 0, err
    }

    createFeatureSetStmt, err := tx.Prepare(`INSERT INTO feature_sets (name, lookup_keys) VALUES (?,?)`)
    if err != nil {
        return 0, err
    }
    defer createFeatureSetStmt.Close()

    result, err := createFeatureSetStmt.Exec(featureSet.Name, strings.Join(featureSet.Keys, ","))
    if err != nil {
        return 0, err
    }

    featureSetId, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    if err = r.createFeatureSetFields(tx, featureSetId, featureSet.Fields); err != nil {
        return 0, err
    }

    if err = tx.Commit(); err != nil {
        return 0, err
    }

    return featureSetId, nil
}

func (r *featuresRepository) Delete(id int64) error {
    defer metrics.Runtime("queries.runtime", []string{"repository:features", "query:delete"})()

    _, err := r.db.ExecPrepared(`DELETE FROM feature_sets WHERE id = ?`, id)
    return err
}

func (r *featuresRepository) scanFeatureSet(rows storage.Scannable) (*FeatureSet, error) {
    featureSet := &FeatureSet{}
    var keys string

    if err := rows.Scan(&featureSet.Id, &featureSet.Name, &keys, &featureSet.CreatedAt, &featureSet.UpdatedAt); err != nil {
        return nil, err
    }
    featureSet.Keys = strings.Split(keys, ",")

    return featureSet, nil
}

func (r *featuresRepository) createFeatureSetFields(tx *sql.Tx, featureSetId int64, fields []*FeatureSetField) error {
    stmt, err := tx.Prepare(fmt.Sprintf(
        `INSERT INTO feature_set_fields (feature_set_id, name, value_type, nullable) VALUES %s`,
        strings.TrimSuffix(strings.Repeat("(?,?,?,?),", len(fields)), ","),
    ))
    if err != nil {
        return err
    }
    defer stmt.Close()

    values := make([]interface{}, 0, len(fields)*4)
    for _, field := range fields {
        values = append(values, featureSetId, field.Name, field.ValueType, field.Nullable)
    }

    _, err = stmt.Exec(values...)
    return err
}

func (r *featuresRepository) queryFeatureSetFields(featureSetId int64) ([]*FeatureSetField, error) {
    rows, err := r.db.Query(`SELECT id, feature_set_id, name, value_type, nullable FROM feature_set_fields WHERE feature_set_id = ?`, featureSetId)
    if err != nil {
        return nil, err
    }

    fields := make([]*FeatureSetField, 0)
    for rows.Next() {
        field := &FeatureSetField{}
        if err = rows.Scan(&field.Id, &field.FeatureSetId, &field.Name, &field.ValueType, &field.Nullable); err != nil {
            return nil, err
        }
        fields = append(fields, field)
    }
    return fields, nil
}

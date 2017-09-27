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

type Model struct {
    Id int64
    Name string `validate:"required"`
    RunnerType string `validate:"required"`
    Features []*ModelFeature `validate:"required,gt=0,dive"`
    PrecomputedFeatures map[int64][]*ModelFeature
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (m *Model) StringId() string {
    return fmt.Sprintf("%d", m.Id)
}

type ModelVersion struct {
    Id int64
    ModelId int64 `validate:"required"`
    Name string `validate:"required"`
    FileName string `validate:"required"`
    IsPrimary bool
    IsShadow bool
    CreatedAt time.Time
}

func (m *ModelVersion) StringId() string {
    return fmt.Sprintf("%d", m.Id)
}

type ModelFeature struct {
    Name string `validate:"required"`
    Alias string
    Required bool
}

type ModelsRepository interface {
    List() ([]*Model, error)
    Find(int64) (*Model, error)
    FindByName(string) (*Model, error)
    Create(*Model) (int64, error)
    Delete(int64) error
    ListVersions(int64) ([]*ModelVersion, error)
    FindVersion(int64) (*ModelVersion, error)
    PrimaryVersion(int64) (*ModelVersion, error)
    CreateVersion(*ModelVersion) (int64, error)
    DeleteVersion(int64) error
    SetPrimaryVersion(int64, int64) error
}

type modelsRepository struct {
    db *storage.Mysql
    validate *validator.Validate
}

func NewModelsRepository(db *storage.Mysql) *modelsRepository {
    return &modelsRepository{
        db: db,
        validate: validator.New(),
    }
}

func (r *modelsRepository) List() ([]*Model, error) {
    defer metrics.Runtime("queries.runtime", []string{"repository:models", "query:list"})()

    rows, err := r.db.Query(`SELECT id, name, runner_type, created_at, updated_at FROM models ORDER BY updated_at DESC`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    models := make([]*Model, 0)
    for rows.Next() {
        model, err := r.scanModel(rows)
        if err != nil {
            return nil, err
        }
        models = append(models, model)
    }

    return models, nil
}

func (r *modelsRepository) Find(id int64) (*Model, error) {
    defer metrics.Runtime("queries.runtime", []string{"repository:models", "query:find"})()

    row, err := r.db.QueryRowPrepared(`SELECT id, name, runner_type, created_at, updated_at FROM models WHERE id = ?`, id)
    if err != nil {
        return nil, err
    }

    model, err := r.scanModel(row)
    if err != nil {
        return nil, err
    }

    if model.Features, err = r.queryModelFeatures(model.Id); err != nil {
        return nil, err
    }

    if model.PrecomputedFeatures, err = r.queryModelPrecomputedFeatures(model.Id); err != nil {
        return nil, err
    }

    return model, nil
}

func (r *modelsRepository) FindByName(name string) (*Model, error) {
    defer metrics.Runtime("queries.runtime", []string{"repository:models", "query:find_by_name"})()

    row, err := r.db.QueryRowPrepared(`SELECT id, name, runner_type, created_at, updated_at FROM models WHERE name = ?`, name)
    if err != nil {
        return nil, err
    }

    model, err := r.scanModel(row)
    if err != nil {
        return nil, err
    }

    if model.Features, err = r.queryModelFeatures(model.Id); err != nil {
        return nil, err
    }

    if model.PrecomputedFeatures, err = r.queryModelPrecomputedFeatures(model.Id); err != nil {
        return nil, err
    }

    return model, nil
}

func (r *modelsRepository) Create(model *Model) (int64, error) {
    if err := r.validate.Struct(model); err != nil {
        return 0, err
    }
    defer metrics.Runtime("queries.runtime", []string{"repository:models", "query:create"})

    tx, err := r.db.Begin()
    if err != nil {
        return 0, err
    }

    createModelStmt, err := tx.Prepare(`INSERT INTO models (name, runner_type) VALUES (?,?)`)
    if err != nil {
        return 0, err
    }
    defer createModelStmt.Close()

    result, err := createModelStmt.Exec(model.Name, model.RunnerType)
    if err != nil {
        return 0, err
    }

    modelId, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    if err = r.createModelFeatures(tx, modelId, model.Features); err != nil {
        return 0, err
    }

    if err = r.createModelPrecomputedFeatures(tx, modelId, model.PrecomputedFeatures); err != nil {
        return 0, err
    }

    if err = tx.Commit(); err != nil {
        return 0, err
    }

    return modelId, nil
}

func (r *modelsRepository) Delete(id int64) error {
    defer metrics.Runtime("queries.runtime", []string{"repository:models", "query:delete"})()

    _, err := r.db.ExecPrepared(`DELETE FROM models WHERE id = ?`, id)

    return err
}

func (r *modelsRepository) scanModel(rows storage.Scannable) (*Model, error) {
    model := &Model{} 
    if err := rows.Scan(&model.Id, &model.Name, &model.RunnerType, &model.CreatedAt, &model.UpdatedAt); err != nil {
        return nil, err
    }
    return model, nil   
}

func (r *modelsRepository) createModelFeatures(tx *sql.Tx, modelId int64, features []*ModelFeature) error {
    stmt, err := tx.Prepare(fmt.Sprintf(
        `INSERT INTO model_features (model_id, name, alias, required) VALUES %s`,
        strings.TrimSuffix(strings.Repeat("(?,?,?,?),", len(features)), ","),
    ))
    if err != nil {
        return err
    }
    defer stmt.Close()

    values := make([]interface{}, 0, len(features)*4)
    for _, feature := range features {
        if feature.Alias == "" {
            feature.Alias = feature.Name
        }
        values = append(values, modelId, feature.Name, feature.Alias, feature.Required)
    }

    _, err = stmt.Exec(values...)
    return err
}

func (r *modelsRepository) createModelPrecomputedFeatures(tx *sql.Tx, modelId int64, features map[int64][]*ModelFeature) error {
    featuresCount, featuresValues := r.precomputedFeaturesValues(modelId, features)
    if featuresCount == 0 {
        return nil
    }

    stmt, err := tx.Prepare(fmt.Sprintf(
        `INSERT INTO model_precomputed_features (model_id, feature_set_id, name, alias, required) VALUES %s`,
        strings.TrimSuffix(strings.Repeat("(?,?,?,?,?),", featuresCount), ","),
    ))
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(featuresValues...)
    return err
}

func (r *modelsRepository) precomputedFeaturesValues(modelId int64, setFeatures map[int64][]*ModelFeature) (int, []interface{}) {
    n, values := 0, make([]interface{}, 0)
    for featureSetId, features := range setFeatures {
        for _, feature := range features {
            n += 1
            if feature.Alias == "" {
                feature.Alias = feature.Name
            }
            values = append(values, modelId, featureSetId, feature.Name, feature.Alias, feature.Required)
        }
    }
    return n, values
}

func (r *modelsRepository) queryModelFeatures(modelId int64) ([]*ModelFeature, error) {
    rows, err := r.db.Query(`SELECT name, alias, required FROM model_features WHERE model_id = ?`, modelId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    features := make([]*ModelFeature, 0)
    for rows.Next() {
        feature := &ModelFeature{}
        if err = rows.Scan(&feature.Name, &feature.Alias, &feature.Required); err != nil {
            return nil, err
        }
        features = append(features, feature)
    }

    return features, nil

}

func (r *modelsRepository) queryModelPrecomputedFeatures(modelId int64) (map[int64][]*ModelFeature, error) {
    rows, err := r.db.Query(`SELECT feature_set_id, name, alias, required FROM model_precomputed_features WHERE model_id = ?`, modelId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    features := make(map[int64][]*ModelFeature, 0)
    for rows.Next() {
        feature := &ModelFeature{}
        var featureSetId int64

        if err = rows.Scan(&featureSetId, &feature.Name, &feature.Alias, &feature.Required); err != nil {
            return nil, err
        }

        if _, exists := features[featureSetId]; !exists {
            features[featureSetId] = make([]*ModelFeature, 0)
        }
        features[featureSetId] = append(features[featureSetId], feature)
    }

    return features, nil
}

func (r *modelsRepository) ListVersions(modelId int64) ([]*ModelVersion, error) {
    defer metrics.Runtime("queries.runtime", []string{"repository:models", "query:list_versions"})()

    rows, err := r.db.QueryPrepared(`SELECT id, model_id, name, file_name, is_primary, is_shadow, created_at FROM model_versions WHERE model_id = ?`, modelId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    versions := make([]*ModelVersion, 0)
    for rows.Next() {
        version, err := r.scanVersion(rows)
        if err != nil {
            return nil, err
        }
        versions = append(versions, version)
    }

    return versions, nil
}

func (r *modelsRepository) FindVersion(id int64) (*ModelVersion, error) {
    defer metrics.Runtime("queries.runtime", []string{"repository:models", "query:find_version"})()

    row, err := r.db.QueryRowPrepared(`SELECT id, model_id, name, file_name, is_primary, is_shadow, created_at FROM model_versions WHERE id = ?`, id)
    if err != nil {
        return nil, err
    }
    return r.scanVersion(row)
}

func (r *modelsRepository) PrimaryVersion(modelId int64) (*ModelVersion, error) {
    defer metrics.Runtime("queries.runtime", []string{"repository:models", "query:primary_version"})()

    row, err := r.db.QueryRowPrepared(`SELECT id, model_id, name, file_name, is_primary, is_shadow, created_at FROM model_versions WHERE model_id = ? AND is_primary = 1`, modelId)
    if err != nil {
        return nil, err
    }
    return r.scanVersion(row)
}

func (r *modelsRepository) CreateVersion(version *ModelVersion) (int64, error) {
    if err := r.validate.Struct(version); err != nil {
        return 0, err
    }
    defer metrics.Runtime("queries.runtime", []string{"repository:models", "query:create_version"})()

    tx, err := r.db.Begin()
    if err != nil {
        return 0, err
    }

    if version.IsPrimary {
        if _, err = tx.Exec(`UPDATE model_versions SET is_primary = 0 WHERE model_id = ?`, version.ModelId); err != nil {
            return 0, err
        }
    }

    createVersionStmt, err := tx.Prepare(`INSERT INTO model_versions (model_id, name, file_name, is_primary, is_shadow) VALUES (?,?,?,?,?)`)
    if err != nil {
        return 0, err
    }
    defer createVersionStmt.Close()

    result, err := createVersionStmt.Exec(version.ModelId, version.Name, version.FileName, version.IsPrimary, version.IsShadow)
    if err != nil {
        return 0, err
    }
    versionId, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    if err = tx.Commit(); err != nil {
        return 0, err
    }

    return versionId, nil
}

func (r *modelsRepository) DeleteVersion(id int64) error {
    defer metrics.Runtime("queries.runtime", []string{"repository:models", "query:delete_version"})()

    version, err := r.FindVersion(id)
    if err != nil {
        return err
    }

    if version.IsPrimary {
        return fmt.Errorf("Cannot delete primary model version id: %d", id)
    }

    _, err = r.db.Exec(`DELETE FROM model_versions WHERE id = ?`, id)
    return err
}

func (r *modelsRepository) SetPrimaryVersion(modelId, versionId int64) error {
    defer metrics.Runtime("queries.runtime", []string{"repository:models", "query:set_primary_version"})()
    
    tx, err := r.db.Begin()
    if err != nil {
        return err
    }

    if _, err := tx.Exec(`UPDATE model_versions SET is_primary = 0 WHERE model_id = ?`, modelId); err != nil {
        return err
    }

    if _, err := tx.Exec(`UPDATE model_versions SET is_primary = 1 WHERE id = ?`, versionId); err != nil {
        return err
    }

    return tx.Commit()
}

func (r *modelsRepository) scanVersion(rows storage.Scannable) (*ModelVersion, error) {
    version := &ModelVersion{}
    if err := rows.Scan(&version.Id, &version.ModelId, &version.Name, &version.FileName, &version.IsPrimary, &version.IsShadow, &version.CreatedAt); err != nil {
        return nil, err
    }
    return version, nil
}

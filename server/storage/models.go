package storage

import (
    "fmt";
    "time";
    "strings";
    "database/sql";

    "github.com/marekgalovic/serving/server/metrics";

    "github.com/satori/go.uuid"
)

type Model struct {
    Uid string
    Name string
    Owner string
    CreatedAt time.Time
    UpdatedAt time.Time
}

type ModelVersion struct {
    Uid string
    Name string
    IsPrimary bool
    IsShadow bool
    RequestFeatures []*ModelFeature
    PrecomputedFeatures map[string][]*ModelFeature
    CreatedAt time.Time
}

type ModelFeature struct {
    Name string
    Required bool
}

type ModelsRepository struct {
    db *Mysql
}

func NewModelsRepository(db *Mysql) *ModelsRepository {
    return &ModelsRepository{
        db: db,
    }
}

func (r *ModelsRepository) List() ([]*Model, error) {
    defer metrics.Runtime("queries.runtime", []string{"repository:models", "query:list"})()

    rows, err := r.db.Query(`SELECT uid, name, owner, created_at, updated_at FROM models ORDER BY updated_at DESC`)
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

func (r *ModelsRepository) Find(uid string) (*Model, error) {
    defer metrics.Runtime("queries.runtime", []string{"repository:models", "query:find"})()

    row, err := r.db.QueryRowPrepared(`SELECT uid, name, owner, created_at, updated_at FROM models WHERE uid = ?`, uid)
    if err != nil {
        return nil, err
    }

    return r.scanModel(row)
}

func (r *ModelsRepository) Create(name, owner string) (*Model, error) {
    defer metrics.Runtime("queries.runtime", []string{"repository:models", "query:create"})

    uid := fmt.Sprintf("%s", uuid.NewV4())

    _, err := r.db.ExecPrepared(`INSERT INTO models (uuid, name, owner) VALUES (?,?,?)`, uid, name, owner)
    if err != nil {
        return nil, err
    }

    return r.Find(uid)
}

func (r *ModelsRepository) Delete(uid string) error {
    defer metrics.Runtime("queries.runtime", []string{"repository:models", "query:delete"})()

    _, err := r.db.ExecPrepared(`DELETE FROM models WHERE uid = ?`, uid)

    return err
}

func (r *ModelsRepository) scanModel(rows Scannable) (*Model, error) {
    model := &Model{} 
    if err := rows.Scan(&model.Uid, &model.Name, &model.Owner, &model.CreatedAt, &model.UpdatedAt); err != nil {
        return nil, err
    }
    return model, nil   
}

func (r *ModelsRepository) ListVersions(modelUid string) ([]*ModelVersion, error) {
    defer metrics.Runtime("queries.runtime", []string{"repository:models", "query:list_versions"})()

    rows, err := r.db.QueryPrepared(`SELECT uid, name, is_primary, is_shadow, request_features, stored_features, created_at FROM model_versions WHERE model_uid = ?`, modelUid)
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

func (r *ModelsRepository) FindVersion(uid string) (*ModelVersion, error) {
    defer metrics.Runtime("queries.runtime", []string{"repository:models", "query:find_version"})()

    row, err := r.db.QueryRowPrepared(`SELECT uid, name, is_primary, is_shadow, request_features, stored_features, created_at FROM model_versions WHERE uid = ?`, uid)
    if err != nil {
        return nil, err
    }
    version, err := r.scanVersion(row)
    if err != nil {
        return nil, err
    }
    version.RequestFeatures, err = r.versionRequestFeatures(version.Uid)
    if err != nil {
        return nil, err
    }
    version.PrecomputedFeatures, err = r.versionPrecomputedFeatures(version.Uid)
    if err != nil {
        return nil, err
    }

    return version, nil
}

func (r *ModelsRepository) PrimaryVersion(modelUid string) (*ModelVersion, error) {
    defer metrics.Runtime("queries.runtime", []string{"repository:models", "query:primary_version"})()

    row, err := r.db.QueryRowPrepared(`SELECT uid, name, is_primary, is_shadow, request_features, stored_features, created_at FROM model_versions WHERE model_uid = ? AND is_primary = 1`, modelUid)
    if err != nil {
        return nil, err
    }
    version, err := r.scanVersion(row)
    if err != nil {
        return nil, err
    }
    version.RequestFeatures, err = r.versionRequestFeatures(version.Uid)
    if err != nil {
        return nil, err
    }
    version.PrecomputedFeatures, err = r.versionPrecomputedFeatures(version.Uid)
    if err != nil {
        return nil, err
    }

    return version, nil
}

func (r *ModelsRepository) CreateVersion(modelUid, name string, isPrimary, isShadow bool, requestFeatures, storedFeatures []string) (*ModelVersion, error) {
    if len(requestFeatures) < 1 {
        return nil, fmt.Errorf("Cannot create model verison with empty request features.")
    }
    defer metrics.Runtime("queries.runtime", []string{"repository:models", "query:create_version"})()
    uid := fmt.Sprintf("%s", uuid.NewV4())

    _, err := r.db.ExecPrepared(
        `INSERT INTO model_versions (uid, model_uid, name, is_shadow, request_features, stored_features) VALUES (?,?,?,?,?,?)`,
        uid, modelUid, name, isShadow, strings.Join(requestFeatures, ","), strings.Join(storedFeatures, ","),
    )
    if err != nil {
        return nil, err
    }

    if isPrimary {
        err := r.SetPrimaryVersion(modelUid, uid)
        if err != nil {
            return nil, fmt.Errorf("Failed to set version as primary. %v", err)
        }
    }

    return r.FindVersion(uid)
}

func (r *ModelsRepository) DeleteVersion(uid string) error {
    defer metrics.Runtime("queries.runtime", []string{"repository:models", "query:delete_version"})()

    version, err := r.FindVersion(uid)
    if err != nil {
        return err
    }

    if version.IsPrimary {
        return fmt.Errorf("Cannot delete primary model version uid: %s", uid)
    }

    _, err = r.db.ExecPrepared(`DELETE FROM model_versions WHERE uid = ?`, uid)
    
    return err
}

func (r *ModelsRepository) SetPrimaryVersion(modelUid, versionUid string) error {
    defer metrics.Runtime("queries.runtime", []string{"repository:models", "query:set_primary_version"})()
    
    tx, err := r.db.Begin()
    if err != nil {
        return err
    }

    unsetPrimaryStmt, err := tx.Prepare(`UPDATE model_versions SET is_primary = 0 WHERE model_uid = ?`)
    if err != nil {
        return err
    }
    defer unsetPrimaryStmt.Close()

    setPrimaryStmt, err := tx.Prepare(`UPDATE model_versions SET is_primary = 1 WHERE uid = ?`)
    if err != nil {
        return err
    }
    defer setPrimaryStmt.Close()
    
    _, err = unsetPrimaryStmt.Exec(modelUid)
    if err != nil {
        return err
    }
    _, err = setPrimaryStmt.Exec(versionUid)
    if err != nil {
        return err
    }

    return tx.Commit()
}

func (r *ModelsRepository) scanVersion(rows Scannable) (*ModelVersion, error) {
    version := &ModelVersion{}
    var requestFeatures sql.NullString
    var storedFeatures sql.NullString

    if err := rows.Scan(&version.Uid, &version.Name, &version.IsPrimary, &version.IsShadow, &requestFeatures, &storedFeatures, &version.CreatedAt); err != nil {
        return nil, err
    }

    _ = requestFeatures
    _ = storedFeatures
    // if requestFeatures.Valid {
    //     version.RequestFeatures = strings.Split(requestFeatures.String, ",")
    // }
    // if storedFeatures.Valid {
    //     version.StoredFeatures = strings.Split(storedFeatures.String, ",")
    // }

    return version, nil
}

func (r *ModelsRepository) versionRequestFeatures(versionUid string) ([]*ModelFeature, error) {
    rows, err := r.db.Query(`SELECT name, required FROM model_version_request_features WHERE model_version_uid = ?`, versionUid)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    features := make([]*ModelFeature, 0)
    for rows.Next() {
        feature := &ModelFeature{}
        if err = rows.Scan(&feature.Name, &feature.Required); err != nil {
            return nil, err
        }
        features = append(features, feature)
    }

    return features, nil

}

func (r *ModelsRepository) versionPrecomputedFeatures(versionUid string) (map[string][]*ModelFeature, error) {
    rows, err := r.db.Query(`SELECT name, required, feature_set_uid FROM model_version_precomputed_features WHERE model_version_uid = ?`, versionUid)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    features := make(map[string][]*ModelFeature, 0)
    for rows.Next() {
        feature := &ModelFeature{}
        var featureSetUid string
        if err = rows.Scan(&feature.Name, &feature.Required, &featureSetUid); err != nil {
            return nil, err
        }

        if _, exists := features[featureSetUid]; !exists {
            features[featureSetUid] = []*ModelFeature{feature}
            continue
        }
        features[featureSetUid] = append(features[featureSetUid], feature)
    }

    return features, nil
}

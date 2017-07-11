package storage

import (
    "fmt";
    "time";
    "strings";
    "database/sql";

    "github.com/satori/go.uuid"

    // log "github.com/Sirupsen/logrus"
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
    RequestFeatures []string
    StoredFeatures []string
    CreatedAt time.Time
    UpdatedAt time.Time
}

type ModelsRepository struct {
    db *Mysql
}

func NewModelsRepository(db *Mysql) *ModelsRepository {
    return &ModelsRepository{
        db: db,
    }
}

func (r *ModelsRepository) Find(uid string) (*Model, error) {
    // Find model
    stmt, err := r.db.Prepare(
        `SELECT uid, name, owner, created_at, updated_at FROM models WHERE models.uid = ?`,
    )
    if err != nil {
        return nil, err
    }
    defer stmt.Close()

    model := &Model{} 
    row := stmt.QueryRow(uid)
    if err = row.Scan(&model.Uid, &model.Name, &model.Owner, &model.CreatedAt, &model.UpdatedAt); err != nil {
        return nil, err
    }
    return model, nil
}

func (r *ModelsRepository) Versions(modelUid string) ([]*ModelVersion, error) {
    stmt, err := r.db.Prepare(
        `SELECT uid, name, is_primary, is_shadow, request_features, stored_features, created_at, updated_at FROM model_versions WHERE model_uid = ?`,
    )
    if err != nil {
        return nil, err
    }
    defer stmt.Close()

    rows, err := stmt.Query(modelUid)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    versions := make([]*ModelVersion, 0)
    for rows.Next() {
        version := &ModelVersion{}
        var requestFeatures sql.NullString
        var storedFeatures sql.NullString

        if err = rows.Scan(&version.Uid, &version.Name, &version.IsPrimary, &version.IsShadow, &requestFeatures, &storedFeatures, &version.CreatedAt, &version.UpdatedAt); err != nil {
            return nil, err
        }
        if requestFeatures.Valid {
            version.RequestFeatures = strings.Split(requestFeatures.String, ",")
        }
        if storedFeatures.Valid {
            version.StoredFeatures = strings.Split(storedFeatures.String, ",")
        }
        versions = append(versions, version)
    }
    return versions, nil
}

func (r *ModelsRepository) PrimaryVersion(modelUid string) (*ModelVersion, error) {
    stmt, err := r.db.Prepare(
        `SELECT uid, name, is_primary, is_shadow, request_features, stored_features, created_at, updated_at FROM model_versions WHERE model_uid = ? AND is_primary = 1`,
    )
    if err != nil {
        return nil, err
    }
    defer stmt.Close()

    version := &ModelVersion{}
    var requestFeatures sql.NullString
    var storedFeatures sql.NullString

    if err = stmt.QueryRow(modelUid).Scan(&version.Uid, &version.Name, &version.IsPrimary, &version.IsShadow, &requestFeatures, &storedFeatures, &version.CreatedAt, &version.UpdatedAt); err != nil {
        return nil, err
    }
    if requestFeatures.Valid {
        version.RequestFeatures = strings.Split(requestFeatures.String, ",")
    }
    if storedFeatures.Valid {
        version.StoredFeatures = strings.Split(storedFeatures.String, ",")
    }

    return version, nil
}

func (r *ModelsRepository) Create(name, owner string) (*Model, error) {
    uid := fmt.Sprintf("%s", uuid.NewV1())

    _, err := r.db.Exec("INSERT INTO models (uuid, name, owner) VALUES (?, ?, ?)", uid, name, owner)
    if err != nil {
        return nil, err
    }

    return r.Find(uid)
}

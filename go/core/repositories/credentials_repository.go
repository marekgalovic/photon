package repositories

import (
    "time";

    "github.com/marekgalovic/photon/go/core/storage";
    "github.com/marekgalovic/photon/go/core/utils";
    "github.com/marekgalovic/photon/go/core/metrics";

    "golang.org/x/crypto/bcrypt";
    "gopkg.in/go-playground/validator.v9";
)

type Credential struct {
    Key string
    Secret string
    Name string `validate:"required"`
    CreatedAt time.Time
}

func (c *Credential) Verify(secret string) error {
    return bcrypt.CompareHashAndPassword([]byte(c.Secret), []byte(secret))
}

type CredentialsRepository interface {
    List() ([]*Credential, error)
    Find(string) (*Credential, error)
    Create(*Credential) (string, string, error)
    Delete(string) error
}

type credentialsRepository struct {
    db *storage.Mysql
    validate *validator.Validate
}

func NewCredentialsRepository(db *storage.Mysql) *credentialsRepository {
    return &credentialsRepository{
        db: db,
        validate: validator.New(),
    }
}

func (r *credentialsRepository) List() ([]*Credential, error) {
    defer metrics.Runtime("queries.runtime", []string{"repository:credentials", "query:list"})()

    rows, err := r.db.Query(`SELECT _key, _secret, name, created_at FROM credentials`)
    if err != nil {
        return nil, err
    }

    credentials := make([]*Credential, 0)
    for rows.Next() {
        credential, err := r.scanCredential(rows)
        if err != nil {
            return nil, err
        }
        credentials = append(credentials, credential)
    }
    
    return credentials, nil
}

func (r *credentialsRepository) Find(key string) (*Credential, error) {
    defer metrics.Runtime("queries.runtime", []string{"repository:credentials", "query:find"})()

    row, err := r.db.QueryRowPrepared(`SELECT _key, _secret, name, created_at FROM credentials WHERE _key = ?`, key)
    if err != nil {
        return nil, err
    }
    return r.scanCredential(row)
}

func (r *credentialsRepository) Create(credential *Credential) (string, string, error) {
    defer metrics.Runtime("queries.runtime", []string{"repository:credentials", "query:create"})()

    key, secret := utils.UuidV4(), utils.UuidV4()
    encryptedSecret, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
    if err != nil {
        return "", "", err
    }

    stmt, err := r.db.Prepare(`INSERT INTO credentials (_key, _secret, name) VALUES (?,?,?,?)`)
    if err != nil {
        return "", "", err
    }
    defer stmt.Close()

    if _, err := stmt.Exec(key, string(encryptedSecret), credential.Name); err != nil {
        return "", "", err
    }
    return key, secret, nil
}

func (r *credentialsRepository) Delete(key string) error {
    defer metrics.Runtime("queries.runtime", []string{"repository:credentials", "query:delete"})()

    _, err := r.db.ExecPrepared(`DELETE FROM credentials WHERE _key = ?`)
    return err
}

func (r *credentialsRepository) scanCredential(row storage.Scannable) (*Credential, error) {
    credential := &Credential{}
    if err := row.Scan(&credential.Key, &credential.Secret, &credential.Name, &credential.CreatedAt); err != nil {
        return nil, err
    }
    return credential, nil
}

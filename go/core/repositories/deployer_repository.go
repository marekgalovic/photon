package repositories

import (
    "fmt";
    "strconv";
    "path/filepath";

    "github.com/marekgalovic/photon/go/core/storage";

    "github.com/samuel/go-zookeeper/zk";
)

type DeployerRepository interface {
    DeployModel(int64) error
    UndeployModel(int64) error
    ModelExistsW(int64) (bool, <-chan zk.Event, error)
    DeployVersion(int64, int64, string) error
    UndeployVersion(int64, int64) error
    ListVersionsW(int64) (map[int64]string, <-chan zk.Event, error)
}

type deployerRepository struct {
    zk *storage.Zookeeper
}

func NewDeployerRepository(zk *storage.Zookeeper) *deployerRepository {
    return &deployerRepository{zk: zk}
}

func (r *deployerRepository) DeployModel(modelId int64) error {
    _, err := r.zk.Create(r.versionsPath(modelId), nil, int32(0), zk.WorldACL(zk.PermAll))
    return err
}

func (r *deployerRepository) UndeployModel(modelId int64) error {
    return r.zk.Delete(r.versionsPath(modelId), -1)
}

func (r *deployerRepository) ModelExistsW(modelId int64) (bool, <-chan zk.Event, error) {
    return r.zk.ExistsW(r.versionsPath(modelId))
}

func (r *deployerRepository) DeployVersion(modelId, versionId int64, fileName string) error {
    versionPath := filepath.Join(r.versionsPath(modelId), fmt.Sprintf("%d", versionId))

    _, err := r.zk.Create(versionPath, fileName, int32(0), zk.WorldACL(zk.PermAll))
    return err
}

func (r *deployerRepository) UndeployVersion(modelId, versionId int64) error {
    versionPath := filepath.Join(r.versionsPath(modelId), fmt.Sprintf("%d", versionId))

    return r.zk.Delete(versionPath, -1)
}

func (r *deployerRepository) ListVersionsW(modelId int64) (map[int64]string, <-chan zk.Event, error) {
    children, event, err := r.zk.ChildrenDataW(r.versionsPath(modelId))
    if err != nil {
        return nil, nil, err
    }

    versions := make(map[int64]string)
    for _, child := range children {
        var fileName string
        if err := child.Scan(&fileName); err != nil {
            return nil, nil, err
        }
        versionId, err := strconv.ParseInt(child.Name, 10, 64)
        if err != nil {
            return nil, nil, err
        }
        versions[versionId] = fileName
    }
    return versions, event, nil
}

func (r *deployerRepository) versionsPath(modelId int64) string {
    return filepath.Join("deployer", fmt.Sprintf("%d", modelId), "versions")
}

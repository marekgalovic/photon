package repositories

import (
    "fmt";
    "regexp";
    "strconv";
    "path/filepath";

    "github.com/marekgalovic/photon/go/core/storage";

    "github.com/samuel/go-zookeeper/zk";
)

var (
    instanceSeqRegexp = regexp.MustCompile(`.*-(\d+)`)
)

type DeployerInstance struct {
    Uid string
    Seq uint64
}

type DeployersRepository struct {
    zk *storage.Zookeeper
}

func NewDeployersRepository(zk *storage.Zookeeper) *DeployersRepository {
    return &DeployersRepository{zk: zk}
}

func (r *DeployersRepository) ListInstances(runnerType string) ([]*DeployerInstance, error) {
    children, err := r.zk.Children(r.instancesPath(runnerType))
    if err != nil {
        return nil, err
    }

    instances := make([]*DeployerInstance, len(children))
    for i, child := range children {
        instance, err := r.childNameToDeployerInstance(child)
        if err != nil {
            return nil, err
        }
        instances[i] = instance
    }
    return instances, nil
}

func (r *DeployersRepository) RegisterInstance(runnerType string) (*DeployerInstance, error) {
    fullPath, err := r.zk.CreateEphemeralSequential(filepath.Join(r.instancesPath(runnerType), "d-"), nil, zk.WorldACL(zk.PermAll))
    if err != nil {
        return nil, err
    }
    
    return r.childNameToDeployerInstance(filepath.Base(fullPath))
}

func (r *DeployersRepository) WatchInstance(runnerType string, uid string) (<-chan zk.Event, error) {
    return r.zk.Watch(filepath.Join(r.instancesPath(runnerType), uid))
}

func (r *DeployersRepository) ListModels(runnerType string) (map[string]int, <-chan zk.Event, error) {
    children, event, err := r.zk.ChildrenDataW(r.modelsPath(runnerType))
    if err != nil {
        return nil, nil, err
    }

    models := make(map[string]int)
    for _, child := range children {
        var config map[string]int
        if err := child.Scan(&config); err != nil {
            return nil, nil, err
        }
        models[child.Name] = config["instances"]
    }

    return models, event, nil
}

func (r *DeployersRepository) DeployModel(runnerType, modelUid string, instancesCount int) error {
    _, err := r.zk.Create(filepath.Join(r.modelsPath(runnerType), modelUid), map[string]int{"instances": instancesCount}, int32(0), zk.WorldACL(zk.PermAll))

    return err
}

func (r *DeployersRepository) UndeployModel(runnerType, modelUid string) error {
    return r.zk.Delete(filepath.Join(r.modelsPath(runnerType), modelUid), -1)
}

func (r *DeployersRepository) instancesPath(runnerType string) string {
    return filepath.Join("deployers", runnerType, "instances")
}

func (r *DeployersRepository) modelsPath(runnerType string) string {
    return filepath.Join("deployers", runnerType, "models")
}

func (r *DeployersRepository) childNameToDeployerInstance(name string) (*DeployerInstance, error) {
    matches := instanceSeqRegexp.FindStringSubmatch(name)
    if len(matches) != 2 {
        return nil, fmt.Errorf("Invalid instance znode name '%s'.", name)
    }

    seq, err := strconv.ParseUint(matches[1], 10, 64)
    if err != nil {
        return nil, err
    } 

    return &DeployerInstance{Uid: name, Seq: seq}, nil  
}

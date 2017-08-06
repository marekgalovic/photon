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
    Seq int
}

type DeployersRepository struct {
    zk *storage.Zookeeper
}

func NewDeployersRepository(zk *storage.Zookeeper) *DeployersRepository {
    return &DeployersRepository{zk: zk}
}

func (r *DeployersRepository) ListInstances(runnerType string) ([]*DeployerInstance, <-chan zk.Event, error) {
    children, event, err := r.zk.ChildrenW(r.instancesPath(runnerType))
    if err != nil {
        return nil, nil, err
    }

    instances := make([]*DeployerInstance, len(children))
    for i, child := range children {
        matches := instanceSeqRegexp.FindStringSubmatch(child)
        if len(matches) != 2 {
            return nil, nil, fmt.Errorf("Invalid instance znode name '%s'.", child)
        }
        seq, err := strconv.Atoi(matches[1])
        if err != nil {
            return nil, nil, err
        }
        instances[i] = &DeployerInstance{Uid: child, Seq: seq}
    }
    
    return instances, event, nil
}

func (r *DeployersRepository) RegisterInstance(runnerType string) error {
    _, err := r.zk.CreateEphemeral(filepath.Join(r.instancesPath(runnerType), "d-"), nil, zk.WorldACL(zk.PermAll))
    
    return err
}

func (r *DeployersRepository) instancesPath(runnerType string) string {
    return filepath.Join("deployers", runnerType, "instances")
}

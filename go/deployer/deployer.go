package deployer

import (
    "fmt";
    "math";
    "time";

    "github.com/marekgalovic/photon/go/core/storage/repositories";

    log "github.com/Sirupsen/logrus"
)

type Deployer struct {
    runnerType string
    deployersRepository *repositories.DeployersRepository
    zookeeperInstance *repositories.DeployerInstance
}

func NewDeployer(runnerType string, deployersRepository *repositories.DeployersRepository) *Deployer {
    return &Deployer{
        runnerType: runnerType,
        deployersRepository: deployersRepository,
    }
}

func (d *Deployer) Run() error {
    instance, err := d.deployersRepository.RegisterInstance(d.runnerType); 
    if err != nil {
        return err
    }
    d.zookeeperInstance = instance

    for {
        instances, err := d.deployersRepository.ListInstances(d.runnerType)
        if err != nil {
            return err
        }
        for _, instance := range instances {
            log.Infof("Instance - uid: %s, seq: %d", instance.Uid, instance.Seq)
        }

        leaderInstance := d.leaderInstance(instances)
        if d.zookeeperInstance.Uid == leaderInstance.Uid {
            log.Info("Local instance is leader.")
            time.Sleep(10 * time.Second)
            log.Info("Done.")
            return nil
        } else {
            event, err := d.deployersRepository.WatchInstance(d.runnerType, leaderInstance.Uid)
            if err != nil {
                return fmt.Errorf("Failed to set watch on leader instance '%s'.", leaderInstance.Uid)
            }
            log.Infof("Set watch on leader instance '%s'.", leaderInstance.Uid)
            <- event
        }
    }

    return nil
}

func (d *Deployer) leaderInstance(instances []*repositories.DeployerInstance) *repositories.DeployerInstance {
    instancesBySeq := make(map[uint64]*repositories.DeployerInstance)
    var minSeq uint64 = math.MaxUint64

    for _, instance := range instances {
        instancesBySeq[instance.Seq] = instance
        if instance.Seq < minSeq {
            minSeq = instance.Seq
        }
    }

    return instancesBySeq[minSeq]
}

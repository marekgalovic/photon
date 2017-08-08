package deployer

import (
    "fmt";
    "math";
    // "time";

    "github.com/marekgalovic/photon/go/core/storage/repositories";

    log "github.com/Sirupsen/logrus"
)

type Deployer struct {
    runnerType string
    deployersRepository *repositories.DeployersRepository
    zookeeperInstance *repositories.DeployerInstance
    errorNotifier chan error
    stopper chan bool
    logger *log.Entry
}

func NewDeployer(runnerType string, deployersRepository *repositories.DeployersRepository) *Deployer {
    return &Deployer{
        runnerType: runnerType,
        deployersRepository: deployersRepository,
        errorNotifier: make(chan error),
        stopper: make(chan bool),
    }
}

func (d *Deployer) Close() {
    close(d.stopper)
}

func (d *Deployer) Run() error {
    instance, err := d.deployersRepository.RegisterInstance(d.runnerType); 
    if err != nil {
        return err
    }
    d.zookeeperInstance = instance
    d.logger = log.WithFields(log.Fields{"deployer_instance_uid": instance.Uid})
    d.logger.Info("Registered instance.")

    go d.watchLeader()
    go d.watchTasks()

    select {
    case err := <- d.errorNotifier:
        return err
    case <- d.stopper:
        return nil
    }
}

func (d *Deployer) watchTasks() error {
    for {
        select {
        case <- d.stopper:
            return nil
        }
    }
    return nil
}

func (d *Deployer) watchLeader() {
    for {
        leader, isLeader, err := d.leaderInstance()
        if err != nil {
            d.errorNotifier <- fmt.Errorf("Failed to get leader instance. %v", err)
            return
        }
        d.logger.Infof("Leader instance uid: %s", leader.Uid)

        if isLeader {
            d.logger.Infof("Instance is leader.")
            go d.watchModels()
            return
        }

        event, err := d.deployersRepository.WatchInstance(d.runnerType, leader.Uid)
        if err != nil {
            d.errorNotifier <- fmt.Errorf("Failed to set watch on leader instance. %v", err)
            return
        }

        select {
        case <- event:
            continue
        case <- d.stopper:
            return
        }
    }
}

func (d *Deployer) watchModels() {
    for {
        models, event, err := d.deployersRepository.ListModels(d.runnerType)
        if err != nil {
            d.errorNotifier <- err
            return
        }

        log.Infof("Loaded models: %d", len(models))

        select {
        case <- event:
            continue
        case <- d.stopper:
            return
        }
    }
}

func (d *Deployer) leaderInstance() (*repositories.DeployerInstance, bool, error) {
    instances, err := d.deployersRepository.ListInstances(d.runnerType)
    if err != nil {
        return nil, false, err
    }

    var leaderInstanceIndex int = 0
    var minSeq uint64 = math.MaxUint64

    for i, instance := range instances {
        if instance.Seq < minSeq {
            minSeq = instance.Seq
            leaderInstanceIndex = i
        }
    }

    return instances[leaderInstanceIndex], instances[leaderInstanceIndex].Uid == d.zookeeperInstance.Uid, nil
}

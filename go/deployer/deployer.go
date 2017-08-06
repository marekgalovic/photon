package deployer

import (
    "github.com/marekgalovic/photon/go/core/storage/repositories";

    log "github.com/Sirupsen/logrus"
)

type Deployer struct {
    runnerType string
    deployersRepository *repositories.DeployersRepository
}

func NewDeployer(runnerType string, deployersRepository *repositories.DeployersRepository) *Deployer {
    return &Deployer{
        runnerType: runnerType,
        deployersRepository: deployersRepository,
    }
}

func (d *Deployer) Run() error {
    if err := d.deployersRepository.RegisterInstance(d.runnerType); err != nil {
        return err
    }

    for {
        instances, event, err := d.deployersRepository.ListInstances(d.runnerType)
        if err != nil {
            return err
        }
        for _, instance := range instances {
            log.Infof("Instance - uid: %s, seq: %d", instance.Uid, instance.Seq)
        }
        <- event
    }

    return nil
}

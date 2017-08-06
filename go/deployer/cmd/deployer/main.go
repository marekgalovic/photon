package main

import (
    "github.com/marekgalovic/photon/go/deployer";
    "github.com/marekgalovic/photon/go/core/storage";
    "github.com/marekgalovic/photon/go/core/storage/repositories";

    log "github.com/Sirupsen/logrus"
)

func main() {
    config, err := deployer.NewConfig()
    if err != nil {
        log.Fatal(err)
    }

    zookeeper, err := storage.NewZookeeper(config.Zookeeper)
    if err != nil {
        log.Fatal(err)
    }
    defer zookeeper.Close()

    deployer := deployer.NewDeployer(config.RunnerType, repositories.NewDeployersRepository(zookeeper))

    if err = deployer.Run(); err != nil {
        log.Fatal(err)
    }
}

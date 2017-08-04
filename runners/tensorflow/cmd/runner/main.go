package main

import (
    "github.com/marekgalovic/photon/runners/tensorflow";
    "github.com/marekgalovic/photon/go/core/storage";
    "github.com/marekgalovic/photon/go/core/storage/repositories";

    log "github.com/Sirupsen/logrus"
)

func main() {
    config, err := runner.NewConfig()
    if err != nil {
        log.Fatal(err)
    }

    zookeeper, err := storage.NewZookeeper(config.Zookeeper)
    if err != nil {
        log.Fatal(err)
    }
    defer zookeeper.Close()

    instancesRepository := repositories.NewInstancesRepository(zookeeper)

    modelManager, err := runner.NewModelManager(config, instancesRepository)
    if err != nil {
        log.Fatal(err)
    }
    defer modelManager.Close()

    modelManager.Watch()

    // log.Info("Running...")
    // for {
    //     uid, err := instancesRepository.Register("model_version_uid_a", config.Address, config.Port)
    //     if err != nil {
    //         log.Fatal(err)
    //     }
    //     time.Sleep(1 * time.Second)
    //     if err = instancesRepository.Unregister("model_version_uid_a", uid); err != nil {
    //         log.Fatal(err)
    //     }
    //     time.Sleep(1 * time.Second)
    // }
}

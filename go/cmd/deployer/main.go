package main

import (
    "github.com/marekgalovic/photon/go/deployer";
    "github.com/marekgalovic/photon/go/core/storage";
    "github.com/marekgalovic/photon/go/core/storage/files";
    "github.com/marekgalovic/photon/go/core/utils";

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

    modelsStore, err := files.NewLocalStorage(files.LocalStorageConfig{Dir: "/Users/marekgalovic/Downloads/uploads"})
    if err != nil {
        log.Fatal(err)
    }

    deployer, err := deployer.NewDeployer(config, zookeeper, modelsStore)
    if err != nil {
        log.Fatal(err)
    }
    defer deployer.Close()

    go func() {
        if err := deployer.Run(); err != nil {
            log.Fatal(err)
        }
    }()

    <- utils.InterruptSignal()
}

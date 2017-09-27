package main

import (
    "github.com/marekgalovic/photon/go/server";
    "github.com/marekgalovic/photon/go/core/storage";
    "github.com/marekgalovic/photon/go/core/storage/files";
    "github.com/marekgalovic/photon/go/core/cluster";
    "github.com/marekgalovic/photon/go/core/storage/features";
    "github.com/marekgalovic/photon/go/core/utils";

    log "github.com/Sirupsen/logrus";
)

func main() {
    config, err := server.NewConfig()
    if err != nil {
        log.Fatal(err)
    }
    config.Print()

    mysql, err := storage.NewMysql(config.Mysql)
    if err != nil {
        log.Fatal(err)
    }
    defer mysql.Close()

    zookeeper, err := storage.NewZookeeper(config.Zookeeper)
    if err != nil {
        log.Fatal(err)
    }
    defer zookeeper.Close()

    kubernetes, err := cluster.NewKubernetes(config.Kubernetes, config.DeploymentTypes)
    if err != nil {
        log.Fatal(err)
    }

    cassandra, err := storage.NewCassandra(config.Cassandra)
    if err != nil {
        log.Fatal(err)
    }
    defer cassandra.Close()
    featuresStore := features.NewCassandraFeaturesStore(cassandra)

    modelsStore, err := files.NewLocalStorage(files.LocalStorageConfig{Dir: "/Users/marekgalovic/Downloads/uploads"})
    if err != nil {
        log.Fatal(err)
    }

    s := server.NewServer(
        config.Server,
        mysql,
        zookeeper,
        kubernetes,
        featuresStore,
        modelsStore,
    )
    if err = s.Serve(); err != nil {
        log.Fatal(err)
    }
    defer s.Stop()

    <- utils.InterruptSignal()
}

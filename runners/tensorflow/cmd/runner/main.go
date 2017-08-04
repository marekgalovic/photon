package main

import (
    "net";

    "github.com/marekgalovic/photon/runners/tensorflow";
    "github.com/marekgalovic/photon/go/core/storage";
    "github.com/marekgalovic/photon/go/core/storage/repositories";
    pb "github.com/marekgalovic/photon/go/core/protos";

    "google.golang.org/grpc";
    log "github.com/Sirupsen/logrus"
)

func main() {
    config, err := runner.NewConfig()
    if err != nil {
        log.Fatal(err)
    }

    listener, err := net.Listen("tcp", config.BindAddress())
    if err != nil {
        log.Fatal(err)
    }
    defer listener.Close()

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
    go modelManager.Watch()

    grpcServer := grpc.NewServer()
    pb.RegisterRunnerServiceServer(grpcServer, runner.NewRunner(modelManager))

    log.Infof("Listening: %s", config.BindAddress())
    grpcServer.Serve(listener)
}

package main

import (
    "net";

    "github.com/marekgalovic/photon/go/server";
    "github.com/marekgalovic/photon/go/core/storage";
    "github.com/marekgalovic/photon/go/core/storage/repositories";
    "github.com/marekgalovic/photon/go/core/storage/features";
    "github.com/marekgalovic/photon/go/server/services";
    pb "github.com/marekgalovic/photon/go/core/protos";

    "google.golang.org/grpc";
    log "github.com/Sirupsen/logrus";
)

func main() {
    config, err := server.NewConfig()
    if err != nil {
        log.Fatalf("Failed to parse config: %v", err)
    }
    config.Print()

    listener, err := net.Listen("tcp", config.BindAddress())
    if err != nil {
        log.Fatal(err)
    }
    defer listener.Close()

    mysql, err := storage.NewMysql(config.Mysql)
    if err != nil {
        log.Fatal(err)
    }
    defer mysql.Close()

    cassandra, err := storage.NewCassandra(config.Cassandra)
    if err != nil {
        log.Fatal(err)
    }
    defer cassandra.Close()

    zookeeper, err := storage.NewZookeeper(config.Zookeeper)
    if err != nil {
        log.Fatal(err)
    }
    defer zookeeper.Close()

    // Stores
    featuresRepository := repositories.NewFeaturesRepository(mysql)
    modelsRepository := repositories.NewModelsRepository(mysql)
    instancesRepository := repositories.NewInstancesRepository(zookeeper)
    featuresStore := features.NewCassandraFeaturesStore(cassandra)

    // Core
    modelResolver := server.NewModelResolver(modelsRepository)
    featuresResolver := server.NewFeaturesResolver(featuresRepository, featuresStore)
    evaluator := server.NewEvaluator(modelResolver, featuresResolver, instancesRepository)
    defer evaluator.Close()

    // Services
    grpcServer := grpc.NewServer()
    pb.RegisterEvaluatorServiceServer(grpcServer, services.NewEvaluatorService(evaluator))
    pb.RegisterModelsServiceServer(grpcServer, services.NewModelsService(modelsRepository))
    pb.RegisterFeaturesServiceServer(grpcServer, services.NewFeaturesService(featuresRepository))

    log.Infof("Listening: %s", config.BindAddress())
    grpcServer.Serve(listener)
}

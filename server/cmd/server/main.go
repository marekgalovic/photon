package main

import (
    "net";

    "github.com/marekgalovic/photon/server";
    "github.com/marekgalovic/photon/server/storage";
    "github.com/marekgalovic/photon/server/storage/repositories";
    "github.com/marekgalovic/photon/server/storage/features";
    "github.com/marekgalovic/photon/server/providers";
    "github.com/marekgalovic/photon/server/services";
    pb "github.com/marekgalovic/photon/server/protos";

    "google.golang.org/grpc";
    log "github.com/Sirupsen/logrus"
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
    featuresStore := features.NewCassandraFeaturesStore(cassandra)

    // Instance provider
    zookeeperProvider := providers.NewZookeeperProvider(zookeeper)

    // Core
    featuresResolver := server.NewFeaturesResolver(featuresRepository, featuresStore)
    modelResolver := server.NewModelResolver(modelsRepository)
    evaluator := server.NewEvaluator(modelResolver, featuresResolver, zookeeperProvider)

    // Services
    grpcServer := grpc.NewServer()
    pb.RegisterEvaluatorServiceServer(grpcServer, services.NewEvaluatorService(evaluator))
    pb.RegisterModelsServiceServer(grpcServer, services.NewModelsService(modelsRepository))
    pb.RegisterFeaturesServiceServer(grpcServer, services.NewFeaturesService(featuresRepository))

    log.Infof("Listening: %s", config.BindAddress())
    grpcServer.Serve(listener)
}

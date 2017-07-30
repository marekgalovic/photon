package main

import (
    "net";

    "github.com/marekgalovic/photon/server";
    "github.com/marekgalovic/photon/server/storage";
    "github.com/marekgalovic/photon/server/storage/repositories";
    "github.com/marekgalovic/photon/server/storage/features";
    "github.com/marekgalovic/photon/server/services";
    pb "github.com/marekgalovic/photon/server/protos";

    "google.golang.org/grpc";
    log "github.com/Sirupsen/logrus"
)

func main() {
    config := server.NewConfig()

    listener, err := net.Listen("tcp", config.BindAddress())
    if err != nil {
        log.Fatal(err)
    }

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

    // Stores
    featuresRepository := repositories.NewFeaturesRepository(mysql)
    modelsRepository := repositories.NewModelsRepository(mysql)
    featuresStore := features.NewFeaturesStore(cassandra)

    // Core
    featuresResolver := server.NewFeaturesResolver(featuresRepository, featuresStore)
    modelResolver := server.NewModelResolver(modelsRepository)
    evaluator := server.NewEvaluator(modelResolver, featuresResolver)

    log.Info(evaluator.Evaluate("f3dbe4f8-68a3-11e7-ab75-0242ac120002", map[string]interface{}{"x1": 1, "x2": 2.83, "x3": "N"}))

    // Services
    grpcServer := grpc.NewServer()
    pb.RegisterEvaluatorServiceServer(grpcServer, services.NewEvaluatorService(evaluator))
    pb.RegisterModelsServiceServer(grpcServer, services.NewModelsService(modelsRepository))
    pb.RegisterFeaturesServiceServer(grpcServer, services.NewFeaturesService(featuresRepository))

    log.Info("Listening ...")
    grpcServer.Serve(listener)
}

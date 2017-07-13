package main

import (
    "net";

    "github.com/marekgalovic/serving/server";
    "github.com/marekgalovic/serving/server/storage";
    "github.com/marekgalovic/serving/server/services";
    pb "github.com/marekgalovic/serving/server/protos";

    "google.golang.org/grpc";
    log "github.com/Sirupsen/logrus"
)

func main() {
    listener, err := net.Listen("tcp", ":5005")
    if err != nil {
        log.Fatal(err)
    }

    mysql, err := storage.NewMysql("root:@tcp(127.0.0.1:3306)/serving_test?parseTime=True")
    if err != nil {
        log.Fatal(err)
    }

    // Stores
    featuresRepository := storage.NewFeaturesRepository(mysql)
    modelsRepository := storage.NewModelsRepository(mysql)

    // Core
    featuresResolver := server.NewFeaturesResolver(featuresRepository)
    modelResolver := server.NewModelResolver(modelsRepository)
    evaluator := server.NewEvaluator(modelResolver, featuresResolver)

    log.Info(modelResolver.GetModel("abcd"))

    // Services
    grpcServer := grpc.NewServer()
    pb.RegisterEvaluatorServiceServer(grpcServer, services.NewEvaluatorService(evaluator))
    pb.RegisterModelsServiceServer(grpcServer, services.NewModelsService(modelsRepository))
    pb.RegisterFeaturesServiceServer(grpcServer, services.NewFeaturesService(featuresRepository))

    log.Info("Listening ...")
    grpcServer.Serve(listener)
}

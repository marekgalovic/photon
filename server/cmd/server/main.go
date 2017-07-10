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

    // Stores
    featuresRepository := storage.NewFeaturesRepository()
    modelsRepository := storage.NewModelsRepository()

    featuresResolver := server.NewFeaturesResolver(featuresRepository)
    modelManager := server.NewModelManager(modelsRepository)
    evaluator := server.NewEvaluator(modelManager, featuresResolver)

    // Services
    grpcServer := grpc.NewServer()
    pb.RegisterEvaluatorServiceServer(grpcServer, services.NewEvaluatorService(evaluator))
    pb.RegisterModelManagerServiceServer(grpcServer, services.NewModelManagerService(modelManager))

    log.Info("Listening ...")
    grpcServer.Serve(listener)
}

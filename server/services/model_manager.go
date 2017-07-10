package services

import (
    "github.com/marekgalovic/serving/server";
    pb "github.com/marekgalovic/serving/server/protos"
)

type ModelManagerService struct {
    modelManager *server.ModelManager
}

func NewModelManagerService(modelManager *server.ModelManager) *ModelManagerService {
    return &ModelManagerService{
        modelManager: modelManager,
    }
}

func (service *ModelManagerService) ListModels(req *pb.ListModelsRequest, stream pb.ModelManagerService_ListModelsServer) error {
    return nil
}

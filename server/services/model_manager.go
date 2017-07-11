package services

import (
    "golang.org/x/net/context";
    
    "github.com/marekgalovic/serving/server/storage";
    pb "github.com/marekgalovic/serving/server/protos"
)

type ModelManagerService struct {
    modelsRepository *storage.ModelsRepository
}

func NewModelManagerService(modelsRepository *storage.ModelsRepository) *ModelManagerService {
    return &ModelManagerService{
        modelsRepository: modelsRepository,
    }
}

func (service *ModelManagerService) ListModels(req *pb.ListModelsRequest, stream pb.ModelManagerService_ListModelsServer) error {
    return nil
}

func (service *ModelManagerService) CreateModel(ctx context.Context, req *pb.CreateModelRequest) (*pb.Model, error) {
    return nil, nil
}

func (service *ModelManagerService) DeleteModel(ctx context.Context, req *pb.DeleteModelRequest) (*pb.EmptyResponse, error) {
    return nil, nil 
}

func (service *ModelManagerService) CreateModelVersion(ctx context.Context, req *pb.CreateModelVersionRequest) (*pb.ModelVersion, error) {
    return nil, nil  
}

func (service *ModelManagerService) DeleteModelVersion(ctx context.Context, req *pb.DeleteModelVersionRequest) (*pb.EmptyResponse, error) {
    return nil, nil  
}

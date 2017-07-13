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

func (service *ModelManagerService) List(req *pb.ListModelsRequest, stream pb.ModelManagerService_ListServer) error {
    models, err := service.modelsRepository.List()
    if err != nil {
        return err
    }

    for _, model := range models {
        if err = stream.Send(service.modelToModelProto(model)); err != nil {
            return err
        }
    }

    return nil
}

func (service *ModelManagerService) Find(ctx context.Context, req *pb.FindModelRequest) (*pb.Model, error) {
    model, err := service.modelsRepository.Find(req.Uid)
    if err != nil {
        return nil, err
    }

    return service.modelToModelProto(model), nil
}

func (service *ModelManagerService) Create(ctx context.Context, req *pb.CreateModelRequest) (*pb.Model, error) {
    model, err := service.modelsRepository.Create(req.Name, req.Owner)
    if err != nil {
        return nil, err
    }

    return service.modelToModelProto(model), nil
}

func (service *ModelManagerService) Delete(ctx context.Context, req *pb.DeleteModelRequest) (*pb.EmptyResponse, error) {
    err := service.modelsRepository.Delete(req.Uid)
    if err != nil {
        return nil, err
    }

    return &pb.EmptyResponse{}, nil
}

func (service *ModelManagerService) modelToModelProto(model *storage.Model) *pb.Model {
    return &pb.Model {
        Uid: model.Uid,
        Name: model.Name,
        Owner: model.Owner,
        CreatedAt: int32(model.CreatedAt.Unix()),
        UpdatedAt: int32(model.UpdatedAt.Unix()), 
    }
}

func (service *ModelManagerService) ListVersions(req *pb.ListVersionsRequest, stream pb.ModelManagerService_ListVersionsServer) error {
    versions, err := service.modelsRepository.ListVersions(req.ModelUid)
    if err != nil {
        return err
    }

    for _, version := range versions {
        if err = stream.Send(service.versionToVersionProto(version)); err != nil {
            return err
        }
    }

    return nil
}

func (service *ModelManagerService) FindVersion(ctx context.Context, req *pb.FindVersionRequest) (*pb.ModelVersion, error) {
    version, err := service.modelsRepository.FindVersion(req.Uid)
    if err != nil {
        return nil, err
    }

    return service.versionToVersionProto(version), nil
}

func (service *ModelManagerService) PrimaryVersion(ctx context.Context, req *pb.PrimaryVersionRequest) (*pb.ModelVersion, error) {
    version, err := service.modelsRepository.PrimaryVersion(req.ModelUid)
    if err != nil {
        return nil, err
    }

    return service.versionToVersionProto(version), nil
}

func (service *ModelManagerService) SetPrimaryVersion(ctx context.Context, req *pb.SetPrimaryVersionRequest) (*pb.EmptyResponse, error) {
    err := service.modelsRepository.SetPrimaryVersion(req.ModelUid, req.Uid)
    if err != nil {
        return nil, err
    }

    return &pb.EmptyResponse{}, nil
}

func (service *ModelManagerService) CreateVersion(ctx context.Context, req *pb.CreateVersionRequest) (*pb.ModelVersion, error) {
    version, err := service.modelsRepository.CreateVersion(req.ModelUid, req.Name, req.IsPrimary, req.IsShadow, req.RequestFeatures, req.StoredFeatures)
    if err != nil {
        return nil, err
    }

    return service.versionToVersionProto(version), nil
}

func (service *ModelManagerService) DeleteVersion(ctx context.Context, req *pb.DeleteVersionRequest) (*pb.EmptyResponse, error) {
    err := service.modelsRepository.DeleteVersion(req.Uid)
    if err != nil {
        return nil, err
    }

    return &pb.EmptyResponse{}, nil
}

func (service *ModelManagerService) versionToVersionProto(version *storage.ModelVersion) *pb.ModelVersion {
    return &pb.ModelVersion {
        Uid: version.Uid,
        Name: version.Name,
        IsPrimary: version.IsPrimary,
        IsShadow: version.IsShadow,
        RequestFeatures: version.RequestFeatures,
        StoredFeatures: version.StoredFeatures,
        CreatedAt: int32(version.CreatedAt.Unix()),
        UpdatedAt: int32(version.UpdatedAt.Unix()),
    }
} 

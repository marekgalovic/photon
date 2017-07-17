package services

import (
    "golang.org/x/net/context";
    
    "github.com/marekgalovic/serving/server/storage";
    pb "github.com/marekgalovic/serving/server/protos"
)

type ModelsService struct {
    modelsRepository *storage.ModelsRepository
}

func NewModelsService(modelsRepository *storage.ModelsRepository) *ModelsService {
    return &ModelsService{
        modelsRepository: modelsRepository,
    }
}

func (service *ModelsService) List(req *pb.ListModelsRequest, stream pb.ModelsService_ListServer) error {
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

func (service *ModelsService) Find(ctx context.Context, req *pb.FindModelRequest) (*pb.Model, error) {
    model, err := service.modelsRepository.Find(req.Uid)
    if err != nil {
        return nil, err
    }

    return service.modelToModelProto(model), nil
}

func (service *ModelsService) Create(ctx context.Context, req *pb.CreateModelRequest) (*pb.Model, error) {
    model, err := service.modelsRepository.Create(req.Name, req.Owner)
    if err != nil {
        return nil, err
    }

    return service.modelToModelProto(model), nil
}

func (service *ModelsService) Delete(ctx context.Context, req *pb.DeleteModelRequest) (*pb.EmptyResponse, error) {
    err := service.modelsRepository.Delete(req.Uid)
    if err != nil {
        return nil, err
    }

    return &pb.EmptyResponse{}, nil
}

func (service *ModelsService) modelToModelProto(model *storage.Model) *pb.Model {
    return &pb.Model {
        Uid: model.Uid,
        Name: model.Name,
        Owner: model.Owner,
        CreatedAt: int32(model.CreatedAt.Unix()),
        UpdatedAt: int32(model.UpdatedAt.Unix()), 
    }
}

func (service *ModelsService) ListVersions(req *pb.ListVersionsRequest, stream pb.ModelsService_ListVersionsServer) error {
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

func (service *ModelsService) FindVersion(ctx context.Context, req *pb.FindVersionRequest) (*pb.ModelVersion, error) {
    version, err := service.modelsRepository.FindVersion(req.Uid)
    if err != nil {
        return nil, err
    }

    return service.versionToVersionProto(version), nil
}

func (service *ModelsService) SetPrimaryVersion(ctx context.Context, req *pb.SetPrimaryVersionRequest) (*pb.EmptyResponse, error) {
    err := service.modelsRepository.SetPrimaryVersion(req.ModelUid, req.Uid)
    if err != nil {
        return nil, err
    }

    return &pb.EmptyResponse{}, nil
}

func (service *ModelsService) CreateVersion(ctx context.Context, req *pb.CreateVersionRequest) (*pb.ModelVersion, error) {
    version, err := service.modelsRepository.CreateVersion(req.ModelUid, req.Name, req.IsPrimary, req.IsShadow, service.requestFeaturesProtoToRequestFeatures(req.RequestFeatures), service.precomputedFeaturesProtoToPrecomputedFeatures(req.PrecomputedFeatures))
    if err != nil {
        return nil, err
    }

    return service.versionToVersionProto(version), nil
}

func (service *ModelsService) DeleteVersion(ctx context.Context, req *pb.DeleteVersionRequest) (*pb.EmptyResponse, error) {
    err := service.modelsRepository.DeleteVersion(req.Uid)
    if err != nil {
        return nil, err
    }

    return &pb.EmptyResponse{}, nil
}

func (service *ModelsService) versionToVersionProto(version *storage.ModelVersion) *pb.ModelVersion {
    return &pb.ModelVersion {
        Uid: version.Uid,
        Name: version.Name,
        IsPrimary: version.IsPrimary,
        IsShadow: version.IsShadow,
        RequestFeatures: service.requestFeaturesToModelFeatureProtos(version.RequestFeatures),
        PrecomputedFeatures: service.precomputedFeaturesToModelFeatureProtos(version.PrecomputedFeatures),
        CreatedAt: int32(version.CreatedAt.Unix()),
    }
} 

func (service *ModelsService) requestFeaturesToModelFeatureProtos(features []*storage.ModelFeature) []*pb.ModelFeature {
    protos := make([]*pb.ModelFeature, 0, len(features))
    for i, feature := range features {
        protos[i] = &pb.ModelFeature{Name: feature.Name, Required: feature.Required}
    }
    return protos
}

func (service *ModelsService) requestFeaturesProtoToRequestFeatures(protos []*pb.ModelFeature) []*storage.ModelFeature {
    features := make([]*storage.ModelFeature, 0, len(protos))
    for i, proto := range protos {
        features[i] = &storage.ModelFeature{Name: proto.Name, Required: proto.Required}
    }
    return features
}

func (service *ModelsService) precomputedFeaturesToModelFeatureProtos(featuresMap map[string][]*storage.ModelFeature) []*pb.ModelFeature {
    protos := make([]*pb.ModelFeature, 0)
    for _, features := range featuresMap {
        for _, feature := range features {
            protos = append(protos, &pb.ModelFeature{Name: feature.Name, Required: feature.Required})
        }
    }
    return protos
}

func (service *ModelsService) precomputedFeaturesProtoToPrecomputedFeatures(protos map[string]*pb.PrecomputedFeaturesSet) map[string][]*storage.ModelFeature {
    features := make(map[string][]*storage.ModelFeature, 0)
    for featureSetUid, precomputedFeaturesSet := range protos {
        if _, exists := features[featureSetUid]; !exists {
            features[featureSetUid] = make([]*storage.ModelFeature, len(precomputedFeaturesSet.Features))
        }
        for _, modelFeatureProto := range precomputedFeaturesSet.Features {
            features[featureSetUid] = append(features[featureSetUid], &storage.ModelFeature{Name: modelFeatureProto.Name, Required: modelFeatureProto.Required})
        }
    }
    return features
}

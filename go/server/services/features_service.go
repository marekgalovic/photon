package services

import (
    "golang.org/x/net/context";
    
    "github.com/marekgalovic/photon/go/core/repositories";
    pb "github.com/marekgalovic/photon/go/core/protos"
)

type FeaturesService struct {
    featuresRepository repositories.FeaturesRepository
}

func NewFeaturesService(featuresRepository repositories.FeaturesRepository) *FeaturesService {
    return &FeaturesService{
        featuresRepository: featuresRepository,
    }
}

func (service *FeaturesService) List(req *pb.EmptyRequest, stream pb.FeaturesService_ListServer) error {
    featureSets, err := service.featuresRepository.List()
    if err != nil {
        return err
    }

    for _, featureSet := range featureSets {
        if err = stream.Send(service.featureSetToFeatureSetProto(featureSet)); err != nil {
            return err
        }
    }

    return nil
}

func (service *FeaturesService) Find(ctx context.Context, req *pb.FindFeatureSetRequest) (*pb.FeatureSet, error) {
    featureSet, err := service.featuresRepository.Find(req.Id)
    if err != nil {
        return nil, err
    }

    return service.featureSetToFeatureSetProto(featureSet), nil
}

func (service *FeaturesService) Create(ctx context.Context, req *pb.CreateFeatureSetRequest) (*pb.CreateFeatureSetResponse, error) {
    featureSet := &repositories.FeatureSet{
        Name: req.Name,
        Keys: req.Keys,
        Fields: service.fieldsProtoToFields(req.Fields),
    }
    featureSetId, err := service.featuresRepository.Create(featureSet)
    if err != nil {
        return nil, err
    }

    return &pb.CreateFeatureSetResponse{Id: featureSetId}, nil
}

func (service *FeaturesService) Delete(ctx context.Context, req *pb.DeleteFeatureSetRequest) (*pb.EmptyResponse, error) {
    err := service.featuresRepository.Delete(req.Id)
    if err != nil {
        return nil, err
    }

    return &pb.EmptyResponse{}, nil
}

func (service *FeaturesService) featureSetToFeatureSetProto(featureSet *repositories.FeatureSet) *pb.FeatureSet {
    return &pb.FeatureSet {
        Id: featureSet.Id,
        Name: featureSet.Name,
        Keys: featureSet.Keys,
        Fields: service.fieldsToFieldsProto(featureSet.Fields),
        CreatedAt: int32(featureSet.CreatedAt.Unix()),
        UpdatedAt: int32(featureSet.UpdatedAt.Unix()),
    }
}

func (service *FeaturesService) fieldsToFieldsProto(fields []*repositories.FeatureSetField) []*pb.FeatureSetField {
    protos := make([]*pb.FeatureSetField, 0, len(fields))
    for i, field := range fields {
        protos[i] = &pb.FeatureSetField{FeatureSetId: field.FeatureSetId, Name: field.Name, ValueType: field.ValueType, Nullable: field.Nullable}
    }
    return protos
}

func (service *FeaturesService) fieldsProtoToFields(protos []*pb.FeatureSetField) []*repositories.FeatureSetField {
    fields := make([]*repositories.FeatureSetField, 0, len(protos))
    for i, proto := range protos {
        fields[i] = &repositories.FeatureSetField{Name: proto.Name, ValueType: proto.ValueType, Nullable: proto.Nullable}
    }
    return fields
}

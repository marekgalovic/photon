package services

import (
    "golang.org/x/net/context";
    
    "github.com/marekgalovic/photon/server/storage";
    pb "github.com/marekgalovic/photon/server/protos"
)

type FeaturesService struct {
    featuresRepository *storage.FeaturesRepository
}

func NewFeaturesService(featuresRepository *storage.FeaturesRepository) *FeaturesService {
    return &FeaturesService{
        featuresRepository: featuresRepository,
    }
}

func (service *FeaturesService) List(req *pb.ListFeatureSetsRequest, stream pb.FeaturesService_ListServer) error {
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
    featureSet, err := service.featuresRepository.Find(req.Uid)
    if err != nil {
        return nil, err
    }

    return service.featureSetToFeatureSetProto(featureSet), nil
}

func (service *FeaturesService) Create(ctx context.Context, req *pb.CreateFeatureSetRequest) (*pb.FeatureSet, error) {
    featureSet, err := service.featuresRepository.Create(req.Name, req.Keys)
    if err != nil {
        return nil, err
    }

    return service.featureSetToFeatureSetProto(featureSet), nil
}

func (service *FeaturesService) Delete(ctx context.Context, req *pb.DeleteFeatureSetRequest) (*pb.EmptyResponse, error) {
    err := service.featuresRepository.Delete(req.Uid)
    if err != nil {
        return nil, err
    }

    return &pb.EmptyResponse{}, nil
}

func (service *FeaturesService) featureSetToFeatureSetProto(featureSet *storage.FeatureSet) *pb.FeatureSet {
    return &pb.FeatureSet {
        Uid: featureSet.Uid,
        Name: featureSet.Name,
        Keys: featureSet.Keys,
        CreatedAt: int32(featureSet.CreatedAt.Unix()),
        UpdatedAt: int32(featureSet.UpdatedAt.Unix()),
    }
}

func (service *FeaturesService) ListSchemas(req *pb.ListFeatureSetSchemasRequest, stream pb.FeaturesService_ListSchemasServer) error {
    schemas, err := service.featuresRepository.ListSchemas(req.FeatureSetUid)
    if err != nil {
        return err
    }

    for _, schema := range schemas {
        if err = stream.Send(service.schemaToSchemaProto(schema)); err != nil {
            return err
        }
    }

    return nil
}

func (service *FeaturesService) FindSchema(ctx context.Context, req *pb.FindFeatureSetSchemaRequest) (*pb.FeatureSetSchema, error) {
    schema, err := service.featuresRepository.FindSchema(req.Uid)
    if err != nil {
        return nil, err
    }

    return service.schemaToSchemaProto(schema), nil
}

func (service *FeaturesService) CreateSchema(ctx context.Context, req *pb.CreateFeatureSetSchemaRequest) (*pb.FeatureSetSchema, error) {
    schema, err := service.featuresRepository.CreateSchema(req.FeatureSetUid, service.schemaFieldsProtoToSchemaFields(req.Fields))
    if err != nil {
        return nil, err
    }

    return service.schemaToSchemaProto(schema), nil
}

func (service *FeaturesService) DeleteSchema(ctx context.Context, req *pb.DeleteFeatureSetSchemaRequest) (*pb.EmptyResponse, error) {
    err := service.featuresRepository.DeleteSchema(req.Uid)
    if err != nil {
        return nil, err
    }

    return &pb.EmptyResponse{}, nil
}

func (service *FeaturesService) schemaToSchemaProto(schema *storage.FeatureSetSchema) *pb.FeatureSetSchema {
    return &pb.FeatureSetSchema {
        Uid: schema.Uid,
        CreatedAt: int32(schema.CreatedAt.Unix()),
    }
}

func (service *FeaturesService) schemaFieldsProtoToSchemaFields(protos []*pb.FeatureSetSchemaField) []*storage.FeatureSetSchemaField {
    fields := make([]*storage.FeatureSetSchemaField, 0, len(protos))
    for i, proto := range protos {
        fields[i] = &storage.FeatureSetSchemaField{Name: proto.Name, ValueType: proto.ValueType, Nullable: proto.Nullable}
    }
    return fields
}

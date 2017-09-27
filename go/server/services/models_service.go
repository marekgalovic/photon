package services

import (
    "io";
    "fmt";
    "golang.org/x/net/context";
    
    "github.com/marekgalovic/photon/go/core/cluster";
    "github.com/marekgalovic/photon/go/core/repositories";
    "github.com/marekgalovic/photon/go/core/storage/files";
    "github.com/marekgalovic/photon/go/core/utils";
    pb "github.com/marekgalovic/photon/go/core/protos";

    // log "github.com/Sirupsen/logrus"
)

type ModelsService struct {
    modelsRepository repositories.ModelsRepository
    deployerRepository repositories.DeployerRepository
    modelsStore files.FilesStore
    kubernetes cluster.Kubernetes
    tmpStore files.FilesStore
}

func NewModelsService(modelsRepository repositories.ModelsRepository, deployerRepository repositories.DeployerRepository, modelsStore files.FilesStore, kubernetes cluster.Kubernetes) *ModelsService {
    tmpStore, err := files.NewLocalStorage(files.LocalStorageConfig{Dir: "/tmp"})
    if err != nil {
        panic(err)
    }
    return &ModelsService{
        modelsRepository: modelsRepository,
        deployerRepository: deployerRepository,
        modelsStore: modelsStore,
        kubernetes: kubernetes,
        tmpStore: tmpStore, 
    }
}

func (service *ModelsService) List(req *pb.EmptyRequest, stream pb.ModelsService_ListServer) error {
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
    model, err := service.modelsRepository.Find(req.Id)
    if err != nil {
        return nil, err
    }

    return service.modelToModelProto(model), nil
}

func (service *ModelsService) Create(ctx context.Context, req *pb.CreateModelRequest) (*pb.CreateModelResponse, error) {
    if !service.kubernetes.HasType(req.RunnerType) {
        return nil, fmt.Errorf("Runner type '%s' does not exists.", req.RunnerType)
    }

    model := &repositories.Model{
        Name: req.Name,
        RunnerType: req.RunnerType,
        Features: service.featuresProtoToFeatures(req.Features),
        PrecomputedFeatures: service.precomputedFeaturesProtoToPrecomputedFeatures(req.PrecomputedFeatures),
    }
    createdModelId, err := service.modelsRepository.Create(model)
    if err != nil {
        return nil, fmt.Errorf("Failed to create model. %v", err)
    }

    if err := service.deployerRepository.DeployModel(createdModelId); err != nil {
        return nil, fmt.Errorf("Failed to deploy model. %v", err)
    }

    deployment := &cluster.ModelDeployment {
        Type: req.RunnerType,
        ModelId: createdModelId,
        Replicas: req.Replicas,
    }
    if err := service.kubernetes.DeployModel(deployment); err != nil {
        return nil, fmt.Errorf("Failed to deploy model. %v", err)
    }

    return &pb.CreateModelResponse{Id: createdModelId}, nil
}

func (service *ModelsService) Delete(ctx context.Context, req *pb.DeleteModelRequest) (*pb.EmptyResponse, error) {
    err := service.modelsRepository.Delete(req.Id)
    if err != nil {
        return nil, err
    }

    if err := service.deployerRepository.UndeployModel(req.Id); err != nil {
        return nil, fmt.Errorf("Failed to undeploy model. %v", err)
    }

    if err := service.kubernetes.UndeployModel(req.Id); err != nil {
        return nil, fmt.Errorf("Failed to undeploy model. %v", err)
    }

    return &pb.EmptyResponse{}, nil
}

func (service *ModelsService) modelToModelProto(model *repositories.Model) *pb.Model {
    return &pb.Model {
        Id: model.Id,
        Name: model.Name,
        Features: service.featuresToFeaturesProto(model.Features),
        PrecomputedFeatures: service.precomputedFeaturesToModelFeatureProtos(model.PrecomputedFeatures),
        CreatedAt: int32(model.CreatedAt.Unix()),
        UpdatedAt: int32(model.UpdatedAt.Unix()), 
    }
}

func (service *ModelsService) ListVersions(req *pb.ListVersionsRequest, stream pb.ModelsService_ListVersionsServer) error {
    versions, err := service.modelsRepository.ListVersions(req.ModelId)
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
    version, err := service.modelsRepository.FindVersion(req.Id)
    if err != nil {
        return nil, err
    }

    return service.versionToVersionProto(version), nil
}

func (service *ModelsService) SetPrimaryVersion(ctx context.Context, req *pb.SetPrimaryVersionRequest) (*pb.EmptyResponse, error) {
    err := service.modelsRepository.SetPrimaryVersion(req.ModelId, req.Id)
    if err != nil {
        return nil, err
    }

    return &pb.EmptyResponse{}, nil
}

func (service *ModelsService) CreateVersion(stream pb.ModelsService_CreateVersionServer) error {
    part, err := stream.Recv()
    if err != nil {
        return err
    }
    versionProto := part.GetVersion()
    if versionProto == nil {
        return fmt.Errorf("Unknown version.")
    }

    tmpName := fmt.Sprintf("%s.xml", utils.UuidV4())
    tmpWriter, err := service.tmpStore.Writer(tmpName)
    if err != nil {
        return err
    }
    defer tmpWriter.Close()

    for {
        part, err := stream.Recv()
        if err == io.EOF {
            return nil
        }
        if err != nil {
            return err
        }
        data := part.GetData()
        if len(data) == 0 {
            break
        }
        tmpWriter.Write(data)

        if err = stream.Send(&pb.CreateVersionResponse{}); err != nil {
            return err
        }
    }

    tmpReader, err := service.tmpStore.Reader(tmpName)
    if err != nil {
        return err
    }
    defer tmpReader.Close()

    permanentWriter, err := service.modelsStore.Writer(tmpName)
    if err != nil {
        return err
    }
    defer permanentWriter.Close()

    if _, err := io.Copy(permanentWriter, tmpReader); err != nil {
        return err
    }

    version := &repositories.ModelVersion{
        ModelId: versionProto.ModelId,
        Name: versionProto.Name,
        FileName: tmpName,
        IsPrimary: versionProto.IsPrimary,
        IsShadow: versionProto.IsShadow,
    }
    createdVersionId, err := service.modelsRepository.CreateVersion(version)
    if err != nil {
        return err
    }

    if err = service.deployerRepository.DeployVersion(version.ModelId, createdVersionId, tmpName); err != nil {
        return err
    }

    if err := stream.Send(&pb.CreateVersionResponse{Id: createdVersionId}); err != nil {
        return err
    }
    return nil
}

func (service *ModelsService) DeleteVersion(ctx context.Context, req *pb.DeleteVersionRequest) (*pb.EmptyResponse, error) {
    version, err := service.modelsRepository.FindVersion(req.Id)
    if err != nil {
        return nil, err
    }

    if err := service.modelsRepository.DeleteVersion(req.Id); err != nil {
        return nil, err
    }

    if err = service.deployerRepository.UndeployVersion(version.ModelId, version.Id); err != nil {
        return nil, err
    }

    return &pb.EmptyResponse{}, nil
}

func (service *ModelsService) versionToVersionProto(version *repositories.ModelVersion) *pb.ModelVersion {
    return &pb.ModelVersion {
        Id: version.Id,
        ModelId: version.ModelId,
        Name: version.Name,
        IsPrimary: version.IsPrimary,
        IsShadow: version.IsShadow,
        CreatedAt: int32(version.CreatedAt.Unix()),
    }
} 

func (service *ModelsService) featuresToFeaturesProto(features []*repositories.ModelFeature) []*pb.ModelFeature {
    protos := make([]*pb.ModelFeature, len(features))
    for i, feature := range features {
        protos[i] = &pb.ModelFeature{Name: feature.Name, Required: feature.Required}
    }
    return protos
}

func (service *ModelsService) featuresProtoToFeatures(protos []*pb.ModelFeature) []*repositories.ModelFeature {
    features := make([]*repositories.ModelFeature, len(protos))
    for i, proto := range protos {
        features[i] = &repositories.ModelFeature{Name: proto.Name, Required: proto.Required}
    }
    return features
}

func (service *ModelsService) precomputedFeaturesToModelFeatureProtos(featuresMap map[int64][]*repositories.ModelFeature) []*pb.ModelFeature {
    protos := make([]*pb.ModelFeature, 0)
    for _, features := range featuresMap {
        for _, feature := range features {
            protos = append(protos, &pb.ModelFeature{Name: feature.Name, Required: feature.Required})
        }
    }
    return protos
}

func (service *ModelsService) precomputedFeaturesProtoToPrecomputedFeatures(protos map[int64]*pb.PrecomputedFeaturesSet) map[int64][]*repositories.ModelFeature {
    features := make(map[int64][]*repositories.ModelFeature, 0)
    for featureSetId, precomputedFeaturesSet := range protos {
        if _, exists := features[featureSetId]; !exists {
            features[featureSetId] = make([]*repositories.ModelFeature, len(precomputedFeaturesSet.Features))
        }
        for _, modelFeatureProto := range precomputedFeaturesSet.Features {
            features[featureSetId] = append(features[featureSetId], &repositories.ModelFeature{Name: modelFeatureProto.Name, Required: modelFeatureProto.Required})
        }
    }
    return features
}

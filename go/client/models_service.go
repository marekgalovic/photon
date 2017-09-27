package photon

import (
    "io";
    "golang.org/x/net/context";

    pb "github.com/marekgalovic/photon/go/core/protos";

    "google.golang.org/grpc";
)

type ModelsService interface {
    ListModels() ([]*pb.Model, error)
    FindModel(int64) (*pb.Model, error)
    CreateModel(string, string, int32, []*pb.ModelFeature, map[int64][]*pb.ModelFeature) (int64, error)
    DeleteModel(int64) error
    ListModelVersions(int64) ([]*pb.ModelVersion, error)
    FindModelVersion(int64) (*pb.ModelVersion, error)
    CreateModelVersion(int64, string, bool, bool, io.Reader) (int64, error)
    DeleteModelVersion(int64) error
}

type modelsService struct {
    client pb.ModelsServiceClient
}

func newModelsService(conn *grpc.ClientConn) *modelsService {
    return &modelsService{
        client: pb.NewModelsServiceClient(conn),
    }
}

func (s *modelsService) ListModels() ([]*pb.Model, error) {
    stream, err := s.client.List(context.Background(), &pb.EmptyRequest{})
    if err != nil {
        return nil, err
    }

    models := make([]*pb.Model, 0)
    for {
        model, err := stream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, err
        }
        models = append(models, model)
    }
    return models, nil
}

func (s *modelsService) FindModel(id int64) (*pb.Model, error) {
    return s.client.Find(context.Background(), &pb.FindModelRequest{Id: id})
}

func (s *modelsService) CreateModel(name, runnerType string, replicas int32, features []*pb.ModelFeature, precomputedFeatures map[int64][]*pb.ModelFeature) (int64, error) {
    precomputedFeatureSets := make(map[int64]*pb.PrecomputedFeaturesSet)
    for featureSetId, features := range precomputedFeatures {
        precomputedFeatureSets[featureSetId] = &pb.PrecomputedFeaturesSet{Features: features}
    }

    response, err := s.client.Create(context.Background(), &pb.CreateModelRequest{
        Name: name,
        RunnerType: runnerType,
        Replicas: replicas,
        Features: features,
        PrecomputedFeatures: precomputedFeatureSets,
    })
    if err != nil {
        return 0, err
    }

    return response.Id, nil
}

func (s *modelsService) DeleteModel(id int64) error {
    _, err := s.client.Delete(context.Background(), &pb.DeleteModelRequest{Id: id})
    return err
}

func (s *modelsService) ListModelVersions(modelId int64) ([]*pb.ModelVersion, error) {
    stream, err := s.client.ListVersions(context.Background(), &pb.ListVersionsRequest{ModelId: modelId})
    if err != nil {
        return nil, err
    }

    versions := make([]*pb.ModelVersion, 0)
    for {
        version, err := stream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, err
        }
        versions = append(versions, version)
    }
    return versions, nil
}

func (s *modelsService) FindModelVersion(id int64) (*pb.ModelVersion, error) {
    return s.client.FindVersion(context.Background(), &pb.FindVersionRequest{Id: id})
}

func (s *modelsService) CreateModelVersion(modelId int64, name string, isPrimary, isShadow bool, modelFile io.Reader) (int64, error) {
    stream, err := s.client.CreateVersion(context.Background())
    if err != nil {
        return 0, err
    }
    defer stream.CloseSend()

    version := &pb.ModelVersion {
        ModelId: modelId,
        Name: name,
        IsPrimary: isPrimary,
        IsShadow: isShadow,
    }
    if err = stream.Send(&pb.CreateVersionRequest{Value: &pb.CreateVersionRequest_Version{Version: version}}); err != nil {
        return 0, err
    }

    for {
        partData := make([]byte, 512 * 1024)
        n, err := modelFile.Read(partData)
        if err == io.EOF {
            if err = stream.Send(&pb.CreateVersionRequest{Value: &pb.CreateVersionRequest_Data{Data: make([]byte, 0)}}); err != nil {
                return 0, err
            }
            break
        }
        if err != nil {
            return 0, err
        }
        if err = stream.Send(&pb.CreateVersionRequest{Value: &pb.CreateVersionRequest_Data{Data: partData[:n]}}); err != nil {
            return 0, err
        }
        if _, err = stream.Recv(); err != nil {
            return 0, err
        }
    }

    response, err := stream.Recv()
    if err != nil {
        return 0, err
    }

    return response.Id, nil
}

func (s *modelsService) DeleteModelVersion(id int64) error {
    _, err := s.client.DeleteVersion(context.Background(), &pb.DeleteVersionRequest{Id: id})
    return err
}

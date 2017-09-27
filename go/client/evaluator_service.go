package photon

import (
    "golang.org/x/net/context";

    pb "github.com/marekgalovic/photon/go/core/protos";
    "github.com/marekgalovic/photon/go/core/utils";
    
    "google.golang.org/grpc";
)

type EvaluatorService interface {
    Evaluate(string, map[string]interface{}) (map[string]interface{}, error)
}

type evaluatorService struct {
    client pb.EvaluatorServiceClient
}

func newEvaluatorService(conn *grpc.ClientConn) *evaluatorService {
    return &evaluatorService {
        client: pb.NewEvaluatorServiceClient(conn),
    }
}

func (s *evaluatorService) Evaluate(modelName string, params map[string]interface{}) (map[string]interface{}, error) {
    valueInterfaces, err := utils.InterfaceMapToValueInterfacePb(params)
    if err != nil {
        return nil, err
    }

    result, err := s.client.Evaluate(context.Background(), &pb.EvaluationRequest{modelName, valueInterfaces})
    if err != nil {
        return nil, err
    }

    return utils.ValueInterfacePbToInterfaceMap(result.Result)
}

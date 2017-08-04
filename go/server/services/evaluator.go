package services

import (
    "golang.org/x/net/context";

    "github.com/marekgalovic/photon/go/server";
    "github.com/marekgalovic/photon/go/core/utils";
    pb "github.com/marekgalovic/photon/go/core/protos"
)

type EvaluatorService struct {
    evaluator *server.Evaluator
}

func NewEvaluatorService(evaluator *server.Evaluator) *EvaluatorService {
    return &EvaluatorService{
        evaluator: evaluator,
    }
}

func (service *EvaluatorService) Evaluate(ctx context.Context, req *pb.EvaluationRequest) (*pb.EvaluationResponse, error) {
    featureInterfaces, err := utils.ValueInterfacePbToInterfaceMap(req.Features)
    if err != nil {
        return nil, err
    }

    resultInterfaces, err := service.evaluator.Evaluate(req.ModelUid, featureInterfaces)
    result, err := utils.InterfaceMapToValueInterfacePb(resultInterfaces)
    if err != nil {
        return nil, err
    }

    return &pb.EvaluationResponse{
        ModelUid: "x000",
        VersionUid: "x001",
        Result: result,
    }, nil
}

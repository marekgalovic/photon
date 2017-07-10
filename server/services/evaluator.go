package services

import (
    "golang.org/x/net/context";

    "github.com/marekgalovic/serving/server";
    pb "github.com/marekgalovic/serving/server/protos"
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
    return nil, nil
}

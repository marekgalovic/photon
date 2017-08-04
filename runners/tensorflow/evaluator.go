package runner

import (
    "golang.org/x/net/context";

    pb "github.com/marekgalovic/photon/go/core/protos";

    log "github.com/Sirupsen/logrus"
)

type Evaluator struct {
    modelManager *ModelManager
}

func NewEvaluator(modelManager *ModelManager) *Evaluator {
    return &Evaluator{
        modelManager: modelManager,
    }
}

func (e *Evaluator) Evaluate(ctx context.Context, req *pb.EvaluationRequest) (*pb.EvaluationResponse, error) {
    log.Info(req)

    return &pb.EvaluationResponse{
        ModelUid: "x000",
        VersionUid: "x001",
    }, nil
}

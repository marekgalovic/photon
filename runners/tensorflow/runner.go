package runner

import (
    "golang.org/x/net/context";

    pb "github.com/marekgalovic/photon/go/core/protos";

    log "github.com/Sirupsen/logrus"
)

type Runner struct {
    modelManager *ModelManager
}

func NewRunner(modelManager *ModelManager) *Runner {
    return &Runner{
        modelManager: modelManager,
    }
}

func (e *Runner) Evaluate(ctx context.Context, req *pb.RunnerEvaluateRequest) (*pb.RunnerEvaluateResponse, error) {
    log.Info(req)

    return &pb.RunnerEvaluateResponse{}, nil
}

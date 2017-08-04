package server

import (
    "golang.org/x/net/context";

    "github.com/marekgalovic/photon/go/core/balancer";
    "github.com/marekgalovic/photon/go/core/storage/repositories";
    pb "github.com/marekgalovic/photon/go/core/protos";

    "google.golang.org/grpc";
    log "github.com/Sirupsen/logrus"
)

type Evaluator struct {
    modelResolver *ModelResolver
    featuresResolver *FeaturesResolver
    instancesRepository *repositories.InstancesRepository
    instancesResolver *balancer.Resolver
}

func NewEvaluator(modelResolver *ModelResolver, featuresResolver *FeaturesResolver, instancesRepository *repositories.InstancesRepository) *Evaluator {
    return &Evaluator{
        modelResolver: modelResolver,
        featuresResolver: featuresResolver,
        instancesRepository: instancesRepository,
        instancesResolver: balancer.NewResolver(instancesRepository),
    }
}

func (e *Evaluator) Evaluate(model_uid string, requestParams map[string]interface{}) (map[string]interface{}, error) {
    _, version, err := e.modelResolver.GetModel(model_uid)
    if err != nil {
        return nil, err
    }

    features, err := e.featuresResolver.Resolve(version, requestParams)
    if err != nil {
        return nil, err
    }

    return e.call(version.Uid, features)
}

func (e *Evaluator) call(versionUid string, features map[string]interface{}) (map[string]interface{}, error) {
    conn, err := grpc.Dial(versionUid, grpc.WithInsecure(), grpc.WithBalancer(grpc.RoundRobin(e.instancesResolver)))
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    client := pb.NewEvaluatorServiceClient(conn)

    result, err := client.Evaluate(context.Background(), &pb.EvaluationRequest{})
    if err != nil {
        return nil, err
    }

    log.Info("gRPC result:", result)

    return nil, nil
}

package server

import (
    "github.com/marekgalovic/photon/go/core/balancer";
    "github.com/marekgalovic/photon/go/core/storage/repositories";

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
    model, version, err := e.modelResolver.GetModel(model_uid)
    if err != nil {
        return nil, err
    }

    features, err := e.featuresResolver.Resolve(version, requestParams)
    if err != nil {
        return nil, err
    }

    log.Info(model)
    log.Info(version)
    log.Info(features)

    conn, err := grpc.Dial(version.Uid, grpc.WithInsecure(), grpc.WithBalancer(grpc.RoundRobin(e.instancesResolver)))
    if err != nil {
        return nil, err
    }
    conn.Close()
    
    return features, nil
}

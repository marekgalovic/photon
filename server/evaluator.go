package server

import (
    "github.com/marekgalovic/photon/server/providers";

    log "github.com/Sirupsen/logrus"
)

type Evaluator struct {
    modelResolver *ModelResolver
    featuresResolver *FeaturesResolver
    instanceProvider providers.InstanceProvider
}

func NewEvaluator(modelResolver *ModelResolver, featuresResolver *FeaturesResolver, instanceProvider providers.InstanceProvider) *Evaluator {
    return &Evaluator{
        modelResolver: modelResolver,
        featuresResolver: featuresResolver,
        instanceProvider: instanceProvider,
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
    instance, err := e.instanceProvider.Get(versionUid)
    if err != nil {
        return nil, err
    }

    log.Info(instance)
    return nil, nil
}

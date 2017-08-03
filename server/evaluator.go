package server

import (
    log "github.com/Sirupsen/logrus"
)

type Evaluator struct {
    modelResolver *ModelResolver
    featuresResolver *FeaturesResolver
    instanceResolver *InstanceResolver
}

func NewEvaluator(modelResolver *ModelResolver, featuresResolver *FeaturesResolver, instanceResolver *InstanceResolver) *Evaluator {
    return &Evaluator{
        modelResolver: modelResolver,
        featuresResolver: featuresResolver,
        instanceResolver: instanceResolver,
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

    return e.call(model.Uid, version.Uid, features)
}

func (e *Evaluator) call(modelUid, versionUid string, features map[string]interface{}) (map[string]interface{}, error) {
    instance, err := e.instanceResolver.Get(versionUid)
    if err != nil {
        return nil, err
    }

    log.Info(instance)
    return nil, nil
}

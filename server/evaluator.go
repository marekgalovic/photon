package server

import (
    log "github.com/Sirupsen/logrus"
)

type Evaluator struct {
    modelResolver *ModelResolver
    featuresResolver *FeaturesResolver
}

func NewEvaluator(modelResolver *ModelResolver, featuresResolver *FeaturesResolver) *Evaluator {
    return &Evaluator{
        modelResolver: modelResolver,
        featuresResolver: featuresResolver,
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

    return nil, nil
}

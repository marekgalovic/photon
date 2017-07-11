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

func (e *Evaluator) Evaluate(model_uid string, features map[string]interface{}) (map[string]interface{}, error) {
    model, version, err := e.modelResolver.GetModel(model_uid)
    if err != nil {
        return nil, err
    }

    log.Info(model)
    log.Info(version)

    // features, err = e.featuresResolver.Resolve(model.PrimaryVersion, features)
    // if err != nil {
    //     return nil, err
    // }

    return nil, nil
}

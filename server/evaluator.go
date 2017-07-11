package server

type Evaluator struct {
    modelManager *ModelManager
    featuresResolver *FeaturesResolver
}

func NewEvaluator(modelManager *ModelManager, featuresResolver *FeaturesResolver) *Evaluator {
    return &Evaluator{
        modelManager: modelManager,
        featuresResolver: featuresResolver,
    }
}

func (e *Evaluator) Evaluate(model_uid string, features map[string]interface{}) (map[string]interface{}, error) {
    model, err := e.modelManager.GetModel(model_uid)
    if err != nil {
        return nil, err
    }

    features, err = e.featuresResolver.Resolve(model.PrimaryVersion, features)
    if err != nil {
        return nil, err
    }

    return features, nil
}

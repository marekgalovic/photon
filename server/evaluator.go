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
    return map[string]interface{}{"score": nil}, nil
}

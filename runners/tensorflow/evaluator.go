package runner

type Evaluator struct {}

func NewEvaluator() *Evaluator {
    return &Evaluator{}
}

func (e *Evaluator) Evaluate(uid string, params map[string]interface{}) (map[string]interface{}, error) {
    return nil, nil
}

func (e *Evaluator) AddModel(uid, modelPath string) error {
    return nil
}

func (e *Evaluator) RemoveModel(uid string) error {
    return nil
}

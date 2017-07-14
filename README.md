# serving
ML model serving service.

## Features
- Serving of many model kinds (sklean, sparkml, tensorfow)
- Static features store
- A/B testing
- Monitoring

models:
    -> MANY: model_versions
        -> MANY: feature_sets{columns}

feature_sets:
    -> MANY: feature_set_schemas

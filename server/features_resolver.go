package server

import (
    "fmt";

    "github.com/marekgalovic/photon/server/storage";
    "github.com/marekgalovic/photon/server/metrics";

    log "github.com/Sirupsen/logrus"
)

type FeaturesResolver struct {
    featuresRepository *storage.FeaturesRepository
}

func NewFeaturesResolver(featuresRepository *storage.FeaturesRepository) *FeaturesResolver {
    return &FeaturesResolver{
        featuresRepository: featuresRepository,
    }
}

func (r *FeaturesResolver) Resolve(version *storage.ModelVersion, requestParams map[string]interface{}) (map[string]interface{}, error) {
    defer metrics.Runtime("features_resolver.resolve.runtime", []string{fmt.Sprintf("model_version_uid:%s", version.Uid)})()

    var precomputedFeatures map[string]interface{}
    var err error
    if len(version.PrecomputedFeatures) > 0 {
        precomputedFeatures, err = r.queryPrecomputedFeatures(version, requestParams)
        if err != nil {
            return nil, err
        }
    }

    return r.merge(version, requestParams, precomputedFeatures)
}

func (r *FeaturesResolver) queryPrecomputedFeatures(version *storage.ModelVersion, requestParams map[string]interface{}) (map[string]interface{}, error) {
    log.Info(version)
    return nil, nil
}

func (r *FeaturesResolver) merge(version *storage.ModelVersion, requestParams map[string]interface{}, precomputedFeatures map[string]interface{}) (map[string]interface{}, error) {
    features := make(map[string]interface{}, 0)

    for _, feature := range version.RequestFeatures {
        value, exists := requestParams[feature.Name]
        if !exists && feature.Required {
            return nil, fmt.Errorf("Required feature '%s' not found in request parameters.", feature.Name)
        }
        features[feature.Name] = value
    }

    for _, features := range version.PrecomputedFeatures {
        for _, feature := range features {
            log.Info(feature)
            // value, exists := precomputedFeatures[feature.Name]
            // if !exists && feature.Required {
            //     return nil, fmt.Errorf("Required feautre '%s' not found in precomputed features.", feature.Name)
            // }
            // features[feature.Name] = value
        }
    }

    return features, nil
}

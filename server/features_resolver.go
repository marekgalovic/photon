package server

import (
    "fmt";
    "time";

    "github.com/marekgalovic/photon/server/storage/repositories";
    "github.com/marekgalovic/photon/server/storage/features";
    "github.com/marekgalovic/photon/server/metrics";

    "github.com/patrickmn/go-cache";
    log "github.com/Sirupsen/logrus"
)

type FeaturesResolver struct {
    featuresRepository *repositories.FeaturesRepository
    featuresStore features.FeaturesStore
    featureSetsCache *cache.Cache
}

type featureSetsCacheEntry struct {
    featureSet *repositories.FeatureSet
    schema *repositories.FeatureSetSchema
}

func NewFeaturesResolver(featuresRepository *repositories.FeaturesRepository, featuresStore features.FeaturesStore) *FeaturesResolver {
    return &FeaturesResolver{
        featuresRepository: featuresRepository,
        featuresStore: featuresStore,
        featureSetsCache: cache.New(30 * time.Second, 1 * time.Minute),
    }
}

func (r *FeaturesResolver) Resolve(version *repositories.ModelVersion, requestParams map[string]interface{}) (map[string]interface{}, error) {
    defer metrics.Runtime("features_resolver.resolve.runtime", []string{fmt.Sprintf("model_version_uid:%s", version.Uid)})()

    if len(version.PrecomputedFeatures) > 0 {
        precomputedFeatures, err := r.queryPrecomputedFeatures(version, requestParams)
        if err != nil {
            return nil, err
        }
        return r.merge(version, requestParams, precomputedFeatures)
    }

    return r.merge(version, requestParams, nil)
}

func (r *FeaturesResolver) queryPrecomputedFeatures(version *repositories.ModelVersion, requestParams map[string]interface{}) (map[string]interface{}, error) {
    log.Info(version)
    for featureSetUid, _ := range version.PrecomputedFeatures {
        featureSet, schema, err := r.getFeatureSet(featureSetUid)
        if err != nil {
            return nil, err
        }
        log.Info(featureSet.Uid, schema.Uid)
    }
    return nil, nil
}

func (r *FeaturesResolver) merge(version *repositories.ModelVersion, requestParams map[string]interface{}, precomputedFeatures map[string]interface{}) (map[string]interface{}, error) {
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

func (r *FeaturesResolver) getFeatureSet(uid string) (*repositories.FeatureSet, *repositories.FeatureSetSchema, error) {
    if cached, exists := r.featureSetsCache.Get(uid); exists {
        entry := cached.(*featureSetsCacheEntry)
        return entry.featureSet, entry.schema, nil
    }

    featureSet, err := r.featuresRepository.Find(uid)
    if err != nil {
        return nil, nil, err
    }

    featureSetSchema, err := r.featuresRepository.LatestSchema(uid)
    if err != nil {
        return nil, nil, err
    }

    r.featureSetsCache.Set(uid, &featureSetsCacheEntry{featureSet: featureSet, schema: featureSetSchema}, cache.DefaultExpiration)
    return featureSet, featureSetSchema, nil
}

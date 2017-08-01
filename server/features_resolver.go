package server

import (
    "fmt";
    "time";

    "github.com/marekgalovic/photon/server/storage/repositories";
    "github.com/marekgalovic/photon/server/storage/features";
    "github.com/marekgalovic/photon/server/metrics";

    "github.com/patrickmn/go-cache";
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
    defer metrics.Runtime("features_resolver.runtime", []string{fmt.Sprintf("model_version_uid:%s", version.Uid), "method:resolve"})()

    requestFeatures, err := r.resolveRequestFeatures(version, requestParams)
    if err != nil {
        return nil, err
    }

    precomputedFeatures, err := r.resolvePrecomputedFeatures(version, requestParams)
    if err != nil {
        return nil, err
    }

    features := make(map[string]interface{})
    for key, value := range requestFeatures {
        features[key] = value
    }
    for key, value := range precomputedFeatures {
        features[key] = value
    }
    return features, nil
}

func (r *FeaturesResolver) resolveRequestFeatures(version *repositories.ModelVersion, requestParams map[string]interface{}) (map[string]interface{}, error) {
    requestFeatures := make(map[string]interface{})

    for _, feature := range version.RequestFeatures {
        value, exists := requestParams[feature.Name]
        if (feature.Required && (!exists || value == nil)) {
            return nil, fmt.Errorf("Required request feature '%s' is missing or null.", feature.Name)
        }
        requestFeatures[feature.Name] = requestParams[feature.Name]
    }

    return requestFeatures, nil
}

func (r *FeaturesResolver) resolvePrecomputedFeatures(version *repositories.ModelVersion, requestParams map[string]interface{}) (map[string]interface{}, error) {
    if len(version.PrecomputedFeatures) == 0 {
        return nil, nil
    }

    queue := make(chan map[string]interface{}, len(version.PrecomputedFeatures))
    finishedNotifier := make(chan bool, 1)
    errNotifier := make(chan error, 1)

    for featureSetUid, features := range version.PrecomputedFeatures {
        go r.queryFeatureSet(featureSetUid, features, requestParams, queue, errNotifier)
    }

    go func() {
        for {
            if len(queue) == cap(queue) {
                finishedNotifier <- true
                return
            }
        }
    }()

    select {
    case <- finishedNotifier:
        precomputedFeatures := make(map[string]interface{})
        for features := range queue {
            for key, value := range features {
                precomputedFeatures[key] = value
            }
        }
        return precomputedFeatures, nil
    case err := <- errNotifier:
        return nil, err
    case <- time.After(1 * time.Second):
        return nil, fmt.Errorf("Timeout while resolving precomputed features.")
    }
}

func (r *FeaturesResolver) queryFeatureSet(featureSetUid string, features []*repositories.ModelFeature, requestParams map[string]interface{}, queue chan map[string]interface{}, errNotifier chan error) {
    featureSet, _, err := r.getFeatureSet(featureSetUid)
    if err != nil {
        errNotifier <- err
        return
    }

    values, err := r.featuresStore.Get(featureSet, requestParams)
    if err != nil {
        errNotifier <- err
        return
    }

    resolvedFeatures := make(map[string]interface{})
    for _, feature := range features {
        value, exists := values[feature.Name]
        if (feature.Required && (!exists || value == nil)) {
            errNotifier <- fmt.Errorf("Required precomputed feature '%s' is missing or null.", feature.Name)
            return
        }
        resolvedFeatures[feature.Name] = values[feature.Name]
    }

    queue <- resolvedFeatures
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

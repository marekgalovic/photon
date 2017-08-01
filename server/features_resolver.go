package server

import (
    "fmt";
    "time";
    "strings";
    "crypto/sha1";

    "github.com/marekgalovic/photon/server/storage";
    "github.com/marekgalovic/photon/server/storage/repositories";
    "github.com/marekgalovic/photon/server/metrics";

    "github.com/patrickmn/go-cache";
)

type FeaturesResolver struct {
    Timeout time.Duration
    featuresRepository *repositories.FeaturesRepository
    featuresStore storage.FeaturesStore
    featureSetsCache *cache.Cache
    featuresCache *cache.Cache
}

type featureSetsCacheEntry struct {
    featureSet *repositories.FeatureSet
    schema *repositories.FeatureSetSchema
}

func NewFeaturesResolver(featuresRepository *repositories.FeaturesRepository, featuresStore storage.FeaturesStore) *FeaturesResolver {
    return &FeaturesResolver{
        Timeout: 100 * time.Millisecond,
        featuresRepository: featuresRepository,
        featuresStore: featuresStore,
        featureSetsCache: cache.New(30 * time.Second, 1 * time.Minute),
        featuresCache: cache.New(10 * time.Second, 30 * time.Second),
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
        requestFeatures[feature.Alias] = requestParams[feature.Name]
    }

    return requestFeatures, nil
}

func (r *FeaturesResolver) resolvePrecomputedFeatures(version *repositories.ModelVersion, requestParams map[string]interface{}) (map[string]interface{}, error) {
    if len(version.PrecomputedFeatures) == 0 {
        return nil, nil
    }

    featureSetsKeys, err := r.featureSetsKeys(version, requestParams)
    if err != nil {
        return nil, err
    }

    cacheKey := r.featuresCacheKey(version, featureSetsKeys)

    if cached, exists := r.featuresCache.Get(cacheKey); exists {
        return cached.(map[string]interface{}), nil
    }

    queue := make(chan map[string]interface{}, len(version.PrecomputedFeatures))
    defer close(queue)
    errNotifier := make(chan error, 1)
    defer close(errNotifier)
    finishedNotifiers := make([]chan bool, 0)
    defer func() {
        for _, notifier := range finishedNotifiers {
            close(notifier)
        }
    }()

    for featureSetUid, features := range version.PrecomputedFeatures {
        finishedNotifier := make(chan bool, 1)
        finishedNotifiers = append(finishedNotifiers, finishedNotifier)
        go r.queryFeatureSet(featureSetUid, features, requestParams, queue, errNotifier, finishedNotifier)
    }

    precomputedFeatures := make(map[string]interface{})
    i := 0
    for {
        select {
        case features := <- queue:
            i += 1
            for key, value := range features {
                precomputedFeatures[key] = value
            }
            if i == len(version.PrecomputedFeatures) {
                r.featuresCache.Set(cacheKey, precomputedFeatures, cache.DefaultExpiration)
                return precomputedFeatures, nil   
            }
        case err := <- errNotifier:
            return nil, err
        case <- time.After(r.Timeout):
            return nil, fmt.Errorf("Timeout while resolving precomputed features.")
        }
    }
}

func (r *FeaturesResolver) featureSetsKeys(version *repositories.ModelVersion, requestParams map[string]interface{}) (map[string]interface{}, error) {
    keys := make(map[string]interface{})
    for featureSetUid, _ := range version.PrecomputedFeatures {
        featureSet, _, err := r.getFeatureSet(featureSetUid)
        if err != nil {
            return nil, err
        }
        for _, key := range featureSet.Keys {
            value, exists := requestParams[key]; 
            if !exists || value == nil {
                return nil, fmt.Errorf("Request param '%s' is missing or null.", key)
            }
            keys[key] = value
        } 
    }
    return keys, nil
}

func (r *FeaturesResolver) featuresCacheKey(version *repositories.ModelVersion, featureSetsKeys map[string]interface{}) string {
    keys := []string{version.Uid}
    for key, value := range featureSetsKeys {
        keys = append(keys, fmt.Sprintf("%s=%v", key, value))
    }
    return fmt.Sprintf("%x", sha1.Sum([]byte(strings.Join(keys, ""))))
}

func (r *FeaturesResolver) queryFeatureSet(featureSetUid string, features []*repositories.ModelFeature, requestParams map[string]interface{}, queue chan map[string]interface{}, errNotifier chan error, finishedNotifier chan bool) {
    featureSet, _, err := r.getFeatureSet(featureSetUid)
    if err != nil {
        select {
        case <- finishedNotifier:
            return
        default:
            errNotifier <- err
            return
        }
    }

    values, err := r.featuresStore.Get(featureSet.Uid, featureSet.Keys, requestParams)
    if err != nil {
        select {
        case <- finishedNotifier:
            return
        default:
            errNotifier <- err
            return
        }
    }

    resolvedFeatures := make(map[string]interface{})
    for _, feature := range features {
        value, exists := values[feature.Name]
        if (feature.Required && (!exists || value == nil)) {
            select {
            case <- finishedNotifier:
                return
            default:
                errNotifier <- fmt.Errorf("Required precomputed feature '%s' is missing or null.", feature.Name)
                return
            }
        }
        resolvedFeatures[feature.Alias] = values[feature.Name]
    }

    select {
    case <- finishedNotifier:
        return
    default:
        queue <- resolvedFeatures    
    }
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

package evaluator

import (
    "fmt";
    "time";
    "strings";
    "crypto/sha1";

    "github.com/marekgalovic/photon/go/core/storage/features";
    "github.com/marekgalovic/photon/go/core/repositories";
    "github.com/marekgalovic/photon/go/core/metrics";

    "github.com/patrickmn/go-cache";
)

type FeaturesResolver interface {
    Resolve(*repositories.Model, map[string]interface{}) (map[string]interface{}, error)
}

type featuresResolver struct {
    Timeout time.Duration
    featuresRepository repositories.FeaturesRepository
    featuresStore features.FeaturesStore
    featureSetsCache *cache.Cache
    featuresCache *cache.Cache
}

func NewFeaturesResolver(featuresRepository repositories.FeaturesRepository, featuresStore features.FeaturesStore) *featuresResolver {
    return &featuresResolver{
        Timeout: 100 * time.Millisecond,
        featuresRepository: featuresRepository,
        featuresStore: featuresStore,
        featureSetsCache: cache.New(30 * time.Second, 1 * time.Minute),
        featuresCache: cache.New(10 * time.Second, 30 * time.Second),
    }
}

func (r *featuresResolver) Resolve(model *repositories.Model, requestParams map[string]interface{}) (map[string]interface{}, error) {
    defer metrics.Runtime("features_resolver.runtime", []string{"method:resolve"})()

    requestFeatures, err := r.resolveRequestFeatures(model, requestParams)
    if err != nil {
        return nil, err
    }

    precomputedFeatures, err := r.resolvePrecomputedFeatures(model, requestParams)
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

func (r *featuresResolver) resolveRequestFeatures(model *repositories.Model, requestParams map[string]interface{}) (map[string]interface{}, error) {
    requestFeatures := make(map[string]interface{})

    for _, feature := range model.Features {
        value, exists := requestParams[feature.Name]
        if (feature.Required && (!exists || value == nil)) {
            return nil, fmt.Errorf("Required request feature '%s' is missing or null.", feature.Name)
        }
        requestFeatures[feature.Alias] = requestParams[feature.Name]
    }

    return requestFeatures, nil
}

func (r *featuresResolver) resolvePrecomputedFeatures(model *repositories.Model, requestParams map[string]interface{}) (map[string]interface{}, error) {
    if len(model.PrecomputedFeatures) == 0 {
        return nil, nil
    }

    featureSetsKeys, err := r.featureSetsKeys(model, requestParams)
    if err != nil {
        return nil, err
    }

    cacheKey := r.featuresCacheKey(model, featureSetsKeys)

    if cached, exists := r.featuresCache.Get(cacheKey); exists {
        return cached.(map[string]interface{}), nil
    }

    queue := make(chan map[string]interface{}, len(model.PrecomputedFeatures))
    defer close(queue)
    errNotifier := make(chan error, 1)
    defer close(errNotifier)
    finishedNotifiers := make([]chan bool, 0)
    defer func() {
        for _, notifier := range finishedNotifiers {
            close(notifier)
        }
    }()

    for featureSetId, features := range model.PrecomputedFeatures {
        finishedNotifier := make(chan bool, 1)
        finishedNotifiers = append(finishedNotifiers, finishedNotifier)
        go r.queryFeatureSet(featureSetId, features, requestParams, queue, errNotifier, finishedNotifier)
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
            if i == len(model.PrecomputedFeatures) {
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

func (r *featuresResolver) featureSetsKeys(model *repositories.Model, requestParams map[string]interface{}) (map[string]interface{}, error) {
    keys := make(map[string]interface{})
    for featureSetUid, _ := range model.PrecomputedFeatures {
        featureSet, err := r.getFeatureSet(featureSetUid)
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

func (r *featuresResolver) featuresCacheKey(model *repositories.Model, featureSetsKeys map[string]interface{}) string {
    keys := []string{model.StringId()}
    for key, value := range featureSetsKeys {
        keys = append(keys, fmt.Sprintf("%s=%v", key, value))
    }
    return fmt.Sprintf("%x", sha1.Sum([]byte(strings.Join(keys, ""))))
}

func (r *featuresResolver) queryFeatureSet(featureSetId int64, features []*repositories.ModelFeature, requestParams map[string]interface{}, queue chan map[string]interface{}, errNotifier chan error, finishedNotifier chan bool) {
    featureSet, err := r.getFeatureSet(featureSetId)
    if err != nil {
        select {
        case <- finishedNotifier:
            return
        default:
            errNotifier <- err
            return
        }
    }

    values, err := r.featuresStore.Get(featureSet.Id, featureSet.Keys, requestParams)
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

func (r *featuresResolver) getFeatureSet(id int64) (*repositories.FeatureSet, error) {
    if cached, exists := r.featureSetsCache.Get(fmt.Sprintf("%d", id)); exists {
        return cached.(*repositories.FeatureSet), nil
    }

    featureSet, err := r.featuresRepository.Find(id)
    if err != nil {
        return nil, err
    }

    r.featureSetsCache.Set(fmt.Sprintf("%d", id), featureSet, cache.DefaultExpiration)
    return featureSet, nil
}

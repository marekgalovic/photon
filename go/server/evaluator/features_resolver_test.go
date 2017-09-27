package evaluator

import (
    "time";
    "testing";

    "github.com/marekgalovic/photon/go/core/repositories";
    "github.com/marekgalovic/photon/go/core/repositories/mock";
    "github.com/marekgalovic/photon/go/core/storage/features/mock";

    "github.com/stretchr/testify/suite";
    "github.com/golang/mock/gomock";
)

type FeaturesResolverTest struct {
    suite.Suite
    mockController *gomock.Controller
    featuresRepository *mock_repositories.MockFeaturesRepository
    featuresStore *mock_features.MockFeaturesStore
    resolver *featuresResolver
    modelWithoutPrecomputedFeatures *repositories.Model
    modelWithPrecomputedFeatures *repositories.Model
    featureSetA *repositories.FeatureSet
    featureSetB *repositories.FeatureSet
}

func TestFeaturesResolver(t *testing.T) {
    suite.Run(t, new(FeaturesResolverTest))
}

func (test *FeaturesResolverTest) SetupSuite() {
    test.modelWithoutPrecomputedFeatures = &repositories.Model{
        Features: []*repositories.ModelFeature{
            {Name: "x1", Alias: "x1", Required: true},
            {Name: "x2", Alias: "x2", Required: false},
        },
    }
    test.modelWithPrecomputedFeatures = &repositories.Model{
        Features: []*repositories.ModelFeature{
            &repositories.ModelFeature{Name: "x1", Alias: "x1", Required: true},
            &repositories.ModelFeature{Name: "x2", Alias: "x2", Required: false},
        },
        PrecomputedFeatures: map[int64][]*repositories.ModelFeature{
            1: []*repositories.ModelFeature{
                {Name: "x1", Alias: "x1_precomputed", Required: true},
                {Name: "x2", Alias: "x2_precomputed", Required: true}, 
            },
            2: []*repositories.ModelFeature{
                {Name: "x3", Alias: "x3", Required: false},
            },
        },
    }
    test.featureSetA = &repositories.FeatureSet{
        Id: 1,
        Keys: []string{"key_a", "key_b"},
        Fields: []*repositories.FeatureSetField{
            {Name: "x1", ValueType: "float", Nullable: false},
            {Name: "x2", ValueType: "float", Nullable: false},
        },
    }
    test.featureSetB = &repositories.FeatureSet{
        Id: 2,
        Keys: []string{"key_c"},
        Fields: []*repositories.FeatureSetField{
            {Name: "x3", ValueType: "float", Nullable: true},
        },
    }
}

func (test *FeaturesResolverTest) SetupTest() {
    test.mockController = gomock.NewController(test.T())
    test.featuresRepository = mock_repositories.NewMockFeaturesRepository(test.mockController)
    test.featuresStore = mock_features.NewMockFeaturesStore(test.mockController)
    test.resolver = NewFeaturesResolver(test.featuresRepository, test.featuresStore)
}

func (test *FeaturesResolverTest) TearDownSuite() {
    test.mockController.Finish()
}

func (test *FeaturesResolverTest) TestResolveWithoutPrecomputedFeatures() {
    features, err := test.resolver.Resolve(test.modelWithoutPrecomputedFeatures, map[string]interface{}{"x1": 1, "x2": 2.5, "foo": "bar"})

    test.Nil(err)
    test.Equal(map[string]interface{}{"x1": 1, "x2": 2.5}, features)
}

func (test *FeaturesResolverTest) TestResolveWithoutPrecomputedFeaturesAndMissingRequiredFeatures() {
    features, err := test.resolver.Resolve(test.modelWithoutPrecomputedFeatures, map[string]interface{}{"x2": 2.5, "foo": "bar"})

    test.NotNil(err)
    test.Nil(features)
}

func (test *FeaturesResolverTest) TestResolveWithoutPrecomputedFeaturesAndMissingOptionalFeatures() {
    features, err := test.resolver.Resolve(test.modelWithoutPrecomputedFeatures, map[string]interface{}{"x1": 1, "foo": "bar"})

    test.Nil(err)
    test.Equal(map[string]interface{}{"x1": 1, "x2": nil}, features)
}

func (test *FeaturesResolverTest) TestResolveWithPrecomputedFeatures() {
    requestParams := map[string]interface{}{"x1": 1, "x2": 2.5, "key_a": 1, "key_b": "foo", "key_c": "bar"}
    test.featuresRepository.EXPECT().Find(test.featureSetA.Id).Return(test.featureSetA, nil)
    test.featuresRepository.EXPECT().Find(test.featureSetB.Id).Return(test.featureSetB, nil)
    test.featuresStore.EXPECT().Get(test.featureSetA.Id, []string{"key_a", "key_b"}, requestParams).Return(map[string]interface{}{"x1": 2.3, "x2": 5.0}, nil)
    test.featuresStore.EXPECT().Get(test.featureSetB.Id, []string{"key_c"}, requestParams).Return(map[string]interface{}{"x3": 7.2}, nil)

    features, err := test.resolver.Resolve(test.modelWithPrecomputedFeatures, requestParams)

    test.Nil(err)
    test.Equal(map[string]interface{}{"x1": 1, "x2": 2.5, "x1_precomputed": 2.3, "x2_precomputed": 5.0, "x3": 7.2}, features)
}

func (test *FeaturesResolverTest) TestResolveWithPrecomputedFeaturesValidatesKeysPresence() {
    requestParams := map[string]interface{}{"x1": 1, "x2": 2.5, "key_a": 1, "key_b": "foo"}
    test.featuresRepository.EXPECT().Find(test.featureSetA.Id).Return(test.featureSetA, nil)
    test.featuresRepository.EXPECT().Find(test.featureSetB.Id).Return(test.featureSetB, nil)

    features, err := test.resolver.Resolve(test.modelWithPrecomputedFeatures, requestParams)

    test.NotNil(err)
    test.Nil(features)
}

func (test *FeaturesResolverTest) TestResolveWithPrecomputedFeaturesAndMissingRequiredPrecomputedFeature() {
    requestParams := map[string]interface{}{"x1": 1, "x2": 2.5, "key_a": 2, "key_b": "foo", "key_c": "bar"}
    test.featuresRepository.EXPECT().Find(test.featureSetA.Id).Return(test.featureSetA, nil)
    test.featuresRepository.EXPECT().Find(test.featureSetB.Id).Return(test.featureSetB, nil)
    test.featuresStore.EXPECT().Get(test.featureSetA.Id, []string{"key_a", "key_b"}, requestParams).Return(map[string]interface{}{"x1": 2.3}, nil)
    test.featuresStore.EXPECT().Get(test.featureSetB.Id, []string{"key_c"}, requestParams).Return(map[string]interface{}{"x3": 7.2}, nil)

    features, err := test.resolver.Resolve(test.modelWithPrecomputedFeatures, requestParams)

    test.NotNil(err)
    test.Nil(features)
}

func (test *FeaturesResolverTest) TestResolverTimeout() {
    requestParams := map[string]interface{}{"x1": 1, "x2": 2.5, "key_a": 2, "key_b": "foo", "key_c": "bar"}
    test.featuresRepository.EXPECT().Find(test.featureSetA.Id).Return(test.featureSetA, nil)
    test.featuresRepository.EXPECT().Find(test.featureSetB.Id).Return(test.featureSetB, nil)
    test.featuresStore.EXPECT().Get(test.featureSetA.Id, []string{"key_a", "key_b"}, requestParams).Do(func(_ ...interface{}) {
        time.Sleep(100 * time.Millisecond)
    })
    test.featuresStore.EXPECT().Get(test.featureSetB.Id, []string{"key_c"}, requestParams).Do(func(_ ...interface{}) {
        time.Sleep(100 * time.Millisecond)
    })

    test.resolver.Timeout = 10 * time.Millisecond
    features, err := test.resolver.Resolve(test.modelWithPrecomputedFeatures, requestParams)

    test.NotNil(err)
    test.Nil(features)
}

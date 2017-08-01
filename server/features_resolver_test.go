package server

import (
    "time";
    "testing";

    "github.com/marekgalovic/photon/server/storage";
    "github.com/marekgalovic/photon/server/storage/repositories";
    "github.com/marekgalovic/photon/server/storage/features";

    "github.com/stretchr/testify/suite";
)

type FeaturesResolverTest struct {
    suite.Suite
    mysql *storage.Mysql
    cassandra *storage.Cassandra
    featuresRepository *repositories.FeaturesRepository
    featuresStore storage.FeaturesStore
    resolver *FeaturesResolver
    modelVersionWithoutPrecomputedFeatures *repositories.ModelVersion
    modelVersionWithPrecomputedFeatures *repositories.ModelVersion
    featureSetA *repositories.FeatureSet
    featureSetB *repositories.FeatureSet
}

func TestFeaturesResolver(t *testing.T) {
    suite.Run(t, new(FeaturesResolverTest))
}

func (test *FeaturesResolverTest) SetupSuite() {
    test.mysql = storage.NewTestMysql()
    test.cassandra = storage.NewTestCassandra()
    test.featuresRepository = repositories.NewFeaturesRepository(test.mysql)
    test.featuresStore = features.NewCassandraFeaturesStore(test.cassandra)

    test.seedDatabase()

    test.modelVersionWithoutPrecomputedFeatures = &repositories.ModelVersion{
        RequestFeatures: []*repositories.ModelFeature{
            &repositories.ModelFeature{Name: "x1", Alias: "x1", Required: true},
            &repositories.ModelFeature{Name: "x2", Alias: "x2", Required: false},
        },
    }
    test.modelVersionWithPrecomputedFeatures = &repositories.ModelVersion{
        RequestFeatures: []*repositories.ModelFeature{
            &repositories.ModelFeature{Name: "x1", Alias: "x1", Required: true},
            &repositories.ModelFeature{Name: "x2", Alias: "x2", Required: false},
        },
        PrecomputedFeatures: map[string][]*repositories.ModelFeature{
            test.featureSetA.Uid: []*repositories.ModelFeature{
                &repositories.ModelFeature{Name: "x1", Alias: "x1_precomputed", Required: true},
                &repositories.ModelFeature{Name: "x2", Alias: "x2_precomputed", Required: true}, 
            },
            test.featureSetB.Uid: []*repositories.ModelFeature{
                &repositories.ModelFeature{Name: "x3", Alias: "x3", Required: false},
            },
        },
    }
}

func (test *FeaturesResolverTest) SetupTest() {
    test.resolver = NewFeaturesResolver(test.featuresRepository, test.featuresStore)
}

func (test *FeaturesResolverTest) TearDownSuite() {
    test.cleanupDatabase()
    test.mysql.Close()
    test.cassandra.Close()
}

func (test *FeaturesResolverTest) seedDatabase() {
    var err error

    test.featureSetA, err = test.featuresRepository.Create("a", []string{"key_a", "key_b"})
    test.Require().Nil(err)

    schemaA, err := test.featuresRepository.CreateSchema(test.featureSetA.Uid, []*repositories.FeatureSetSchemaField{
        &repositories.FeatureSetSchemaField{Name: "x1", ValueType: "float", Nullable: false},
        &repositories.FeatureSetSchemaField{Name: "x2", ValueType: "float", Nullable: false},
    })
    test.Require().Nil(err)

    test.featureSetB, err = test.featuresRepository.Create("b", []string{"key_c"})
    test.Require().Nil(err)

    schemaB, err := test.featuresRepository.CreateSchema(test.featureSetB.Uid, []*repositories.FeatureSetSchemaField{
        &repositories.FeatureSetSchemaField{Name: "x3", ValueType: "float", Nullable: true},
    })
    test.Require().Nil(err)

    err = test.featuresStore.CreateFeatureSet(test.featureSetA.Uid, test.featureSetA.Keys)
    test.Require().Nil(err)

    err = test.featuresStore.CreateFeatureSet(test.featureSetB.Uid, test.featureSetB.Keys)
    test.Require().Nil(err)

    err = test.featuresStore.Insert(test.featureSetA.Uid, schemaA.Uid, test.featureSetA.Keys, map[string]interface{}{"key_a": 1, "key_b": "foo", "x1": 2.3, "x2": 5.0})
    test.Require().Nil(err)

    err = test.featuresStore.Insert(test.featureSetA.Uid, schemaA.Uid, test.featureSetA.Keys, map[string]interface{}{"key_a": 2, "key_b": "foo", "x1": 2.3})
    test.Require().Nil(err)

    err = test.featuresStore.Insert(test.featureSetB.Uid, schemaB.Uid, test.featureSetB.Keys, map[string]interface{}{"key_c": "bar", "x3": 7.2})
    test.Require().Nil(err)
}

func (test *FeaturesResolverTest) cleanupDatabase() {
    storage.CleanupMysql(test.mysql, "models", "feature_sets")
    test.featuresStore.DeleteFeatureSet(test.featureSetA.Uid)
    test.featuresStore.DeleteFeatureSet(test.featureSetB.Uid)
}

func (test *FeaturesResolverTest) TestResolveWithoutPrecomputedFeatures() {
    features, err := test.resolver.Resolve(test.modelVersionWithoutPrecomputedFeatures, map[string]interface{}{"x1": 1, "x2": 2.5, "foo": "bar"})

    test.Nil(err)
    test.Equal(map[string]interface{}{"x1": 1, "x2": 2.5}, features)
}

func (test *FeaturesResolverTest) TestResolveWithoutPrecomputedFeaturesAndMissingRequiredFeatures() {
    features, err := test.resolver.Resolve(test.modelVersionWithoutPrecomputedFeatures, map[string]interface{}{"x2": 2.5, "foo": "bar"})

    test.NotNil(err)
    test.Nil(features)
}

func (test *FeaturesResolverTest) TestResolveWithoutPrecomputedFeaturesAndMissingOptionalFeatures() {
    features, err := test.resolver.Resolve(test.modelVersionWithoutPrecomputedFeatures, map[string]interface{}{"x1": 1, "foo": "bar"})

    test.Nil(err)
    test.Equal(map[string]interface{}{"x1": 1, "x2": nil}, features)
}

func (test *FeaturesResolverTest) TestResolveWithPrecomputedFeatures() {
    features, err := test.resolver.Resolve(test.modelVersionWithPrecomputedFeatures, map[string]interface{}{"x1": 1, "x2": 2.5, "key_a": 1, "key_b": "foo", "key_c": "bar"})

    test.Nil(err)
    test.Equal(map[string]interface{}{"x1": 1, "x2": 2.5, "x1_precomputed": 2.3, "x2_precomputed": 5.0, "x3": 7.2}, features)
}

func (test *FeaturesResolverTest) TestResolveWithPrecomputedFeaturesValidatesKeysPresence() {
    features, err := test.resolver.Resolve(test.modelVersionWithPrecomputedFeatures, map[string]interface{}{"x1": 1, "x2": 2.5, "key_a": 1, "key_b": "foo"})

    test.NotNil(err)
    test.Nil(features)
}

func (test *FeaturesResolverTest) TestResolveWithPrecomputedFeaturesAndMissingRequiredPrecomputedFeature() {
    features, err := test.resolver.Resolve(test.modelVersionWithPrecomputedFeatures, map[string]interface{}{"x1": 1, "x2": 2.5, "key_a": 2, "key_b": "foo", "key_c": "bar"})

    test.NotNil(err)
    test.Nil(features)
}

func (test *FeaturesResolverTest) TestResolverTimeout() {
    test.resolver.Timeout = 1 * time.Millisecond
    features, err := test.resolver.Resolve(test.modelVersionWithPrecomputedFeatures, map[string]interface{}{"x1": 1, "x2": 2.5, "key_a": 2, "key_b": "foo", "key_c": "bar"})

    test.NotNil(err)
    test.Nil(features)
}

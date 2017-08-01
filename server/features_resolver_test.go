package server

import (
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
    test.resolver = NewFeaturesResolver(test.featuresRepository, test.featuresStore)

    test.seedDatabase()

    test.modelVersionWithoutPrecomputedFeatures = &repositories.ModelVersion{
        RequestFeatures: []*repositories.ModelFeature{
            &repositories.ModelFeature{Name: "x1", Required: true},
            &repositories.ModelFeature{Name: "x2", Required: false},
        },
    }
    test.modelVersionWithPrecomputedFeatures = &repositories.ModelVersion{
        RequestFeatures: []*repositories.ModelFeature{
            &repositories.ModelFeature{Name: "x1", Required: true},
            &repositories.ModelFeature{Name: "x2", Required: false},
        },
        PrecomputedFeatures: map[string][]*repositories.ModelFeature{
            test.featureSetA.Uid: []*repositories.ModelFeature{
                &repositories.ModelFeature{Name: "x3", Required: true},
                &repositories.ModelFeature{Name: "x4", Required: true}, 
            },
            test.featureSetB.Uid: []*repositories.ModelFeature{
                &repositories.ModelFeature{Name: "x5", Required: false},
            },
        },
    }
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

    _, err = test.featuresRepository.CreateSchema(test.featureSetA.Uid, []*repositories.FeatureSetSchemaField{
        &repositories.FeatureSetSchemaField{Name: "x3", ValueType: "float", Nullable: false},
        &repositories.FeatureSetSchemaField{Name: "x4", ValueType: "float", Nullable: false},
    })
    test.Require().Nil(err)

    test.featureSetB, err = test.featuresRepository.Create("b", []string{"key_c"})
    test.Require().Nil(err)

    _, err = test.featuresRepository.CreateSchema(test.featureSetB.Uid, []*repositories.FeatureSetSchemaField{
        &repositories.FeatureSetSchemaField{Name: "x5", ValueType: "float", Nullable: true},
    })
    test.Require().Nil(err)

    err = test.featuresStore.CreateFeatureSet(test.featureSetA.Uid, test.featureSetA.Keys)
    test.Require().Nil(err)

    err = test.featuresStore.CreateFeatureSet(test.featureSetB.Uid, test.featureSetB.Keys)
    test.Require().Nil(err)
}

func (test *FeaturesResolverTest) cleanupDatabase() {
    storage.CleanupMysql(test.mysql, "models", "feature_sets")
    test.featuresStore.DeleteFeatureSet(test.featureSetA.Uid)
    test.featuresStore.DeleteFeatureSet(test.featureSetB.Uid)
}

func (test *FeaturesResolverTest) TestResolveWithoutPrecomputedFeatures() {
    features, err := test.resolver.Resolve(test.modelVersionWithoutPrecomputedFeatures, map[string]interface{}{"x1": 1, "x2": 2.5, "foo": "bar"})

    test.Require().Nil(err)
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

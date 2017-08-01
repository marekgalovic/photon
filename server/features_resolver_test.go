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
    mysqlDb *storage.Mysql
    cassandraDb *storage.Cassandra
    resolver *FeaturesResolver
}

func TestFeaturesResolver(t *testing.T) {
    suite.Run(t, new(FeaturesResolverTest))
}

func (test *FeaturesResolverTest) SetupSuite() {
    test.mysqlDb = storage.NewTestMysql()
    test.cassandraDb = storage.NewTestCassandra()
    test.resolver = NewFeaturesResolver(repositories.NewFeaturesRepository(test.mysqlDb), features.NewCassandraFeaturesStore(test.cassandraDb))
}

func (test *FeaturesResolverTest) TearDownSuite() {
    test.mysqlDb.Close()
    test.cassandraDb.Close()
}

func (test *FeaturesResolverTest) TestResolveWithoutPrecomputedFeatures() {
    modelVersion := &repositories.ModelVersion{
        RequestFeatures: []*repositories.ModelFeature{
            &repositories.ModelFeature{Name: "x1", Required: true},
            &repositories.ModelFeature{Name: "x2", Required: false},
        },
    }

    features, err := test.resolver.Resolve(modelVersion, map[string]interface{}{"x1": 1, "x2": 2.5, "foo": "bar"})

    test.Require().Nil(err)
    test.Equal(map[string]interface{}{"x1": 1, "x2": 2.5}, features)
}

func (test *FeaturesResolverTest) TestResolveWithoutPrecomputedFeaturesAndMissingRequiredFeatures() {
    modelVersion := &repositories.ModelVersion{
        RequestFeatures: []*repositories.ModelFeature{
            &repositories.ModelFeature{Name: "x1", Required: true},
            &repositories.ModelFeature{Name: "x2", Required: false},
        },
    }

    features, err := test.resolver.Resolve(modelVersion, map[string]interface{}{"x2": 2.5, "foo": "bar"})

    test.NotNil(err)
    test.Nil(features)
}

func (test *FeaturesResolverTest) TestResolveWithoutPrecomputedFeaturesAndMissingOptionalFeatures() {
    modelVersion := &repositories.ModelVersion{
        RequestFeatures: []*repositories.ModelFeature{
            &repositories.ModelFeature{Name: "x1", Required: true},
            &repositories.ModelFeature{Name: "x2", Required: false},
        },
    }

    features, err := test.resolver.Resolve(modelVersion, map[string]interface{}{"x1": 1, "foo": "bar"})

    test.Nil(err)
    test.Equal(map[string]interface{}{"x1": 1, "x2": nil}, features)
}

package repositories

import (
    "testing";

    "github.com/marekgalovic/photon/go/core/storage";

    "github.com/stretchr/testify/suite";
)

type FeaturesRepositoryTest struct {
    suite.Suite
    db *storage.Mysql
    repository *featuresRepository
    featureSetId int64
}

func TestFeaturesRepository(t *testing.T) {
    suite.Run(t, new(FeaturesRepositoryTest))
}

func (test *FeaturesRepositoryTest) SetupTest() {
    test.db = storage.NewTestMysql()
    storage.CleanupMysql(test.db, "models", "feature_sets")

    test.repository = NewFeaturesRepository(test.db)
    test.seedDatabase()
}

func (test *FeaturesRepositoryTest) TearDownTest() {
    test.db.Close()
}

func (test *FeaturesRepositoryTest) seedDatabase() {
    result, err := test.db.Exec(
        `INSERT INTO feature_sets (name, lookup_keys) VALUES (?,?)`,
        "precomputed_features", "resource_id,email",
    )
    test.Require().Nil(err, "Failed to seed database.")

    test.featureSetId, err = result.LastInsertId()
    test.Require().Nil(err, "Failed to seed database.")

    _, err = test.db.Exec(
        `INSERT INTO feature_set_fields (feature_set_id, name, value_type, nullable) VALUES (?,?,?,?),(?,?,?,?),(?,?,?,?)`,
        test.featureSetId, "x1", "integer", false,
        test.featureSetId, "x2", "double", true,
        test.featureSetId, "x3", "boolean", true,
    )
    test.Require().Nil(err)
}

func (test *FeaturesRepositoryTest) TestList() {
    featureSets, err := test.repository.List()

    test.Require().Nil(err)
    test.Equal(1, len(featureSets))
}

func (test *FeaturesRepositoryTest) TestFind() {
    set, err := test.repository.Find(test.featureSetId)

    test.Require().Nil(err)
    test.Equal("precomputed_features", set.Name)
    test.Equal([]string{"resource_id", "email"}, set.Keys)
    test.Equal(3, len(set.Fields))
}

func (test *FeaturesRepositoryTest) TestCreate() {
    storage.AssertCountChanged(test.db, "feature_sets", 1, func() {
        storage.AssertCountChanged(test.db, "feature_set_fields", 2, func() {
            featureSet := &FeatureSet{
                Name: "customer_email_features",
                Keys: []string{"email"},
                Fields: []*FeatureSetField{
                    {Name: "x1", ValueType: "integer"},
                    {Name: "x2", ValueType: "float", Nullable: true},
                },
            }
            _, err := test.repository.Create(featureSet)
            test.Nil(err)
        })
    })
}

func (test *FeaturesRepositoryTest) TestDelete() {
    storage.AssertCountChanged(test.db, "feature_sets", -1, func() {
        storage.AssertCountChanged(test.db, "feature_set_fields", -3, func() {
            err := test.repository.Delete(test.featureSetId)
            test.Nil(err)
        })
    })
}

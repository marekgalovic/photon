package repositories

import (
    "testing";

    "github.com/marekgalovic/photon/server/storage";

    "github.com/stretchr/testify/suite";
)

type FeaturesRepositoryTest struct {
    suite.Suite
    db *storage.Mysql
    repository *FeaturesRepository
}

func TestFeaturesRepository(t *testing.T) {
    suite.Run(t, new(FeaturesRepositoryTest))
}

func (test *FeaturesRepositoryTest) SetupTest() {
    test.db = storage.NewTestMysql()
    storage.CleanupMysql(test.db)

    test.repository = NewFeaturesRepository(test.db)
    test.seedDatabase()
}

func (test *FeaturesRepositoryTest) TearDownTest() {
    test.db.Close()
}

func (test *FeaturesRepositoryTest) seedDatabase() {
    _, err := test.db.Exec(
        `INSERT INTO feature_sets (uid, name, lookup_keys) VALUES (?,?,?),(?,?,?)`,
        "test-feature-set-uid", "precomputed_features", "resource_id",
        "test-feature-set-uid2", "historic_features", "resource_id,email",
    )
    test.Require().Nil(err)

    _, err = test.db.Exec(
        `INSERT INTO feature_set_schemas (uid, feature_set_uid) VALUES (?,?),(?,?),(?,?)`,
        "test-feature-set-schema-uid", "test-feature-set-uid",
        "test-feature-set-schema-uid2", "test-feature-set-uid",
        "test-feature-set-schema-uid3", "test-feature-set-uid2",
    )
    test.Require().Nil(err)

    _, err = test.db.Exec(
        `INSERT INTO feature_set_schema_fields (feature_set_schema_uid, name, value_type, nullable) VALUES (?,?,?,?),(?,?,?,?)`,
        "test-feature-set-schema-uid", "x1", "float", false,
        "test-feature-set-schema-uid", "x2", "integer", true,
    )
    test.Require().Nil(err)
}

func (test *FeaturesRepositoryTest) TestList() {
    featureSets, err := test.repository.List()

    test.Require().Nil(err)
    test.Equal(2, len(featureSets))
}

func (test *FeaturesRepositoryTest) TestFind() {
    set, err := test.repository.Find("test-feature-set-uid2")

    test.Require().Nil(err)
    test.Equal("historic_features", set.Name)
    test.Equal([]string{"resource_id", "email"}, set.Keys)
}

func (test *FeaturesRepositoryTest) TestCreate() {
    storage.AssertCountChanged(test.db, "feature_sets", 1, func() {
        _, err := test.repository.Create("customer_email_features", []string{"email"})
        test.Nil(err)
    })
}

func (test *FeaturesRepositoryTest) TestDelete() {
    storage.AssertCountChanged(test.db, "feature_sets", -1, func() {
        storage.AssertCountChanged(test.db, "feature_set_schemas", -2, func() {
            storage.AssertCountChanged(test.db, "feature_set_schema_fields", -2, func() {
                err := test.repository.Delete("test-feature-set-uid")
                test.Nil(err)
            })
        })
    })
}

func (test *FeaturesRepositoryTest) TestListSchemas() {
    schemas, err := test.repository.ListSchemas("test-feature-set-uid")
    test.Require().Nil(err)

    test.Equal(2, len(schemas))
}

func (test *FeaturesRepositoryTest) TestFindSchema() {
    schema, err := test.repository.FindSchema("test-feature-set-schema-uid")
    test.Require().Nil(err)

    test.Equal(2, len(schema.Fields))
    test.Equal(&FeatureSetSchemaField{"x1", "float", false}, schema.Fields[0])
    test.Equal(&FeatureSetSchemaField{"x2", "integer", true}, schema.Fields[1])
}

func (test *FeaturesRepositoryTest) TestLatestSchema() {
    schema, err := test.repository.LatestSchema("test-feature-set-uid")

    test.Require().Nil(err)
    test.Equal("test-feature-set-schema-uid2", schema.Uid)
}

func (test *FeaturesRepositoryTest) TestCreateSchema() {
    storage.AssertCountChanged(test.db, "feature_set_schemas", 1, func() {
        storage.AssertCountChanged(test.db, "feature_set_schema_fields", 2, func() {
            _, err := test.repository.CreateSchema("test-feature-set-uid", []*FeatureSetSchemaField{&FeatureSetSchemaField{"x1", "float", true}, &FeatureSetSchemaField{"x2", "boolean", false}})
            test.Nil(err)
        })
    })
}

func (test *FeaturesRepositoryTest) TestDeleteSchema() {
    storage.AssertCountChanged(test.db, "feature_set_schemas", -1, func() {
        storage.AssertCountChanged(test.db, "feature_set_schema_fields", -2, func() {
            err := test.repository.DeleteSchema("test-feature-set-schema-uid")
            test.Nil(err)
        })
    })
}

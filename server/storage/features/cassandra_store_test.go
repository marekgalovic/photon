package features

import (
    "testing";

    "github.com/marekgalovic/photon/server/storage";

    "github.com/stretchr/testify/suite";
)

type CassandraFeaturesStoreTest struct {
    suite.Suite
    db *storage.Cassandra
    store *CassandraFeaturesStore
}

func TestCassandraFeaturesStore(t *testing.T) {
    suite.Run(t, new(CassandraFeaturesStoreTest))
}

func (test *CassandraFeaturesStoreTest) SetupSuite() {
    test.db = storage.NewTestCassandra()
    test.store = NewCassandraFeaturesStore(test.db)
}

func (test *CassandraFeaturesStoreTest) TearDownSuite() {
    test.db.Close()
}

func (test *CassandraFeaturesStoreTest) SetupTest() {
    test.seedDatabase()
}

func (test *CassandraFeaturesStoreTest) seedDatabase() {
    err := test.db.Query(`DROP TABLE IF EXISTS set_test_features_uid`).Exec();
    test.Require().Nil(err)

    err = test.db.Query(`DROP TABLE IF EXISTS set_test_features_uid2`).Exec();
    test.Require().Nil(err)

    err = test.db.Query(`CREATE TABLE set_test_features_uid (schema_uid VARCHAR, key_a VARCHAR, key_b VARCHAR, data TEXT, PRIMARY KEY (key_a, key_b))`).Exec()
    test.Require().Nil(err)

    err = test.db.Query(`INSERT INTO set_test_features_uid (schema_uid, key_a, key_b, data) VALUES ('test-schema-uid', '1', 'foo@bar.com', '{"x1": "foo", "x2": 1, "x3": 2.5}')`).Exec()
    test.Require().Nil(err)
}
 
func (test *CassandraFeaturesStoreTest) TestGet() {
    features, err := test.store.Get("test-features-uid", []string{"key_a", "key_b"}, map[string]interface{}{"key_a": 1, "key_b": "foo@bar.com"})
    test.Require().Nil(err)

    test.Require().NotNil(features)
    test.Equal("foo", features["x1"])
    test.Equal(float64(1), features["x2"])
    test.Equal(2.5, features["x3"])
}

func (test *CassandraFeaturesStoreTest) TestInsert() {
    storage.AssertCountChanged(test.db, "set_test_features_uid", 1, func() {
        err := test.store.Insert(
            "test-features-uid", "3e53a72b-70ba-4255-8d89-f96de7c1c6b9", []string{"key_a", "key_b"},
            map[string]interface{}{"key_a": 2, "key_b": "foo@baz.com", "x1": "baz", "x2": 2, "x3": 2.6},
        )
        test.Nil(err)
    })
}

func (test *CassandraFeaturesStoreTest) TestCreateFeatureSet() {
    err := test.store.CreateFeatureSet("test-features-uid2", []string{"key_c"})

    test.Nil(err)
}

func (test *CassandraFeaturesStoreTest) TestDeleteFeatureSet() {
    err := test.store.DeleteFeatureSet("test-features-uid")

    test.Nil(err)
}

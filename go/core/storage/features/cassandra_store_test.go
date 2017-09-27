package features

import (
    "testing";

    "github.com/marekgalovic/photon/go/core/storage";

    "github.com/stretchr/testify/suite";
)

type CassandraFeaturesStoreTest struct {
    suite.Suite
    db *storage.Cassandra
    store *cassandraFeaturesStore
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
    err := test.db.Query(`DROP TABLE IF EXISTS feature_set_1`).Exec();
    test.Require().Nil(err)

    err = test.db.Query(`DROP TABLE IF EXISTS feature_set_2`).Exec();
    test.Require().Nil(err)

    err = test.db.Query(`CREATE TABLE feature_set_1 (key_a VARCHAR, key_b VARCHAR, data TEXT, PRIMARY KEY (key_a, key_b))`).Exec()
    test.Require().Nil(err)

    err = test.db.Query(`INSERT INTO feature_set_1 (key_a, key_b, data) VALUES ('1', 'foo@bar.com', '{"x1": "foo", "x2": 1, "x3": 2.5}')`).Exec()
    test.Require().Nil(err)
}
 
func (test *CassandraFeaturesStoreTest) TestGet() {
    features, err := test.store.Get(1, []string{"key_a", "key_b"}, map[string]interface{}{"key_a": 1, "key_b": "foo@bar.com"})
    test.Require().Nil(err)

    test.Require().NotNil(features)
    test.Equal("foo", features["x1"])
    test.Equal(float64(1), features["x2"])
    test.Equal(2.5, features["x3"])
}

func (test *CassandraFeaturesStoreTest) TestInsert() {
    storage.AssertCountChanged(test.db, "feature_set_1", 1, func() {
        err := test.store.Insert(
            1, []string{"key_a", "key_b"},
            map[string]interface{}{"key_a": 2, "key_b": "foo@baz.com", "x1": "baz", "x2": 2, "x3": 2.6},
        )
        test.Require().Nil(err)
    })
}

func (test *CassandraFeaturesStoreTest) TestCreateFeatureSet() {
    err := test.store.CreateFeatureSet(2, []string{"key_c"})

    test.Nil(err)
}

func (test *CassandraFeaturesStoreTest) TestDeleteFeatureSet() {
    err := test.store.DeleteFeatureSet(1)

    test.Nil(err)
}

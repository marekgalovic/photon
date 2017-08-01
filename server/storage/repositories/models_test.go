package repositories

import (
    "testing";

    "github.com/marekgalovic/photon/server/storage";

    "github.com/stretchr/testify/suite";
)

type ModelsRepositoryTest struct {
    suite.Suite
    db *storage.Mysql
    repository *ModelsRepository
}

func TestModelsRepository(t *testing.T) {
    suite.Run(t, new(ModelsRepositoryTest))
}

func (test *ModelsRepositoryTest) SetupTest() {
    test.db = storage.NewTestMysql()
    storage.CleanupMysql(test.db, "models", "feature_sets")

    test.repository = NewModelsRepository(test.db)
    test.seedDatabase()
}

func (test *ModelsRepositoryTest) TearDownTest() {
    test.db.Close()
}

func (test *ModelsRepositoryTest) seedDatabase() {
    _, err := test.db.Exec(
        `INSERT INTO models (uid, name, owner) VALUES (?,?,?),(?,?,?)`,
        "test-model-uid", "test_model", "foo@bar.com",
        "test-model-uid2", "test_model_2", "foo@baz.com",
    )
    test.Require().Nil(err, "Failed to seed models")

    _, err = test.db.Exec(
        `INSERT INTO model_versions (uid, model_uid, name, is_shadow, is_primary) VALUES (?,?,?,?,?),(?,?,?,?,?)`,
        "test-version-uid", "test-model-uid", "test_version", false, true,
        "test-version-uid2", "test-model-uid", "test_version_2", false, false,
    )
    test.Require().Nil(err)

    _, err = test.db.Exec(
        `INSERT INTO model_version_request_features (model_version_uid, name, required) VALUES (?,?,?),(?,?,?)`,
        "test-version-uid", "x1", true,
        "test-version-uid", "x2", false,
    )
    test.Require().Nil(err)

    _, err = test.db.Exec(`INSERT INTO feature_sets (uid, name, lookup_keys) VALUES (?,?,?)`, "test-feature-set-uid", "precomputed_features", "resource_id")
    test.Require().Nil(err)

    _, err = test.db.Exec(
        `INSERT INTO model_version_precomputed_features (model_version_uid, feature_set_uid, name, required) VALUES (?,?,?,?),(?,?,?,?)`,
        "test-version-uid", "test-feature-set-uid", "x3", true,
        "test-version-uid", "test-feature-set-uid", "x4", false,
    )
    test.Require().Nil(err)
}

func (test *ModelsRepositoryTest) TestList() {
    models, err := test.repository.List()
    test.Require().Nil(err)

    test.Equal(2, len(models))
}

func (test *ModelsRepositoryTest) TestFind() {
    model, err := test.repository.Find("test-model-uid")
    test.Require().Nil(err)

    test.Equal("test-model-uid", model.Uid)
    test.Equal("test_model", model.Name)
    test.Equal("foo@bar.com", model.Owner)
}

func (test *ModelsRepositoryTest) TestCreate() {
    storage.AssertCountChanged(test.db, "models", 1, func() {
        _, err := test.repository.Create("test_model_2", "foo@baz.com")
        test.Nil(err)
    })
}

func (test *ModelsRepositoryTest) TestDelete() {
    storage.AssertCountChanged(test.db, "models", -1, func() {
        err := test.repository.Delete("test-model-uid")
        test.Nil(err)
    })
}

func (test *ModelsRepositoryTest) TestListVersions() {
    versions, err := test.repository.ListVersions("test-model-uid")
    test.Require().Nil(err)

    test.Equal(2, len(versions))
}

func (test *ModelsRepositoryTest) TestFindVersion() {
    version, err := test.repository.FindVersion("test-version-uid")
    test.Require().Nil(err)

    test.Equal("test-version-uid", version.Uid)
    test.Equal("test_version", version.Name)
    test.Equal(false, version.IsShadow)
    test.Equal(true, version.IsPrimary)
    test.Equal(2, len(version.RequestFeatures))
    test.Equal(1, len(version.PrecomputedFeatures))
    test.Equal(2, len(version.PrecomputedFeatures["test-feature-set-uid"]))
}

func (test *ModelsRepositoryTest) TestPrimaryVersion() {
    version, err := test.repository.PrimaryVersion("test-model-uid")
    test.Require().Nil(err)

    test.Equal("test-version-uid", version.Uid)
    test.Equal(true, version.IsPrimary)
}

func (test *ModelsRepositoryTest) TestPrimaryVersionForModelWithNoVersions() {
    version, err := test.repository.PrimaryVersion("test-model-uid2")

    test.NotNil(err)
    test.Nil(version)
}

func (test *ModelsRepositoryTest) TestCreateVersion() {
    storage.AssertCountChanged(test.db, "model_versions", 1, func() {
        storage.AssertCountChanged(test.db, "model_version_request_features", 1, func() {
            version, err := test.repository.CreateVersion("test-model-uid", "test_version_3", false, false, []*ModelFeature{&ModelFeature{"x1", true}}, map[string][]*ModelFeature{"test-feature-set-uid": []*ModelFeature{&ModelFeature{"x2", false}}})
            test.Require().Nil(err)

            test.NotNil("test_version_3", version.Name)
            test.Equal(false, version.IsShadow)
            test.Equal(false, version.IsPrimary)
            test.Equal(1, len(version.RequestFeatures))
            test.Equal(1, len(version.PrecomputedFeatures))
            test.Equal(&ModelFeature{"x1", true}, version.RequestFeatures[0])
            test.Equal(&ModelFeature{"x2", false}, version.PrecomputedFeatures["test-feature-set-uid"][0])
        })
    })
}

func (test *ModelsRepositoryTest) TestCreatePrimaryVersion() {
    oldPrimaryVersion, err := test.repository.PrimaryVersion("test-model-uid")
    test.Require().Nil(err)

    version, err := test.repository.CreateVersion("test-model-uid", "test_version_3", true, false, []*ModelFeature{&ModelFeature{"x1", false}}, map[string][]*ModelFeature{})
    test.Require().Nil(err)

    newPrimaryVersion, err := test.repository.PrimaryVersion("test-model-uid")
    test.Require().Nil(err)

    test.Equal(version.Uid, newPrimaryVersion.Uid)
    test.NotEqual(oldPrimaryVersion.Uid, newPrimaryVersion.Uid)
}

func (test *ModelsRepositoryTest) TestDeleteVersion() {
    storage.AssertCountChanged(test.db, "model_versions", -1, func() {
        err := test.repository.DeleteVersion("test-version-uid2")
        test.Nil(err)
    })
}

func (test *ModelsRepositoryTest) TestDeleteVersionWhenVersionIsPrimary() {
    storage.AssertCountChanged(test.db, "model_versions", 0, func() {
        err := test.repository.DeleteVersion("test-version-uid")
        test.NotNil(err)
    })
}

func (test *ModelsRepositoryTest) TestSetPrimaryVersion() {
    primaryVersion, err := test.repository.PrimaryVersion("test-model-uid")
    test.Require().Nil(err)
    test.Equal("test-version-uid", primaryVersion.Uid)

    err = test.repository.SetPrimaryVersion("test-model-uid", "test-version-uid2")
    test.Nil(err)

    primaryVersion, err = test.repository.PrimaryVersion("test-model-uid")
    test.Require().Nil(err)
    test.Equal("test-version-uid2", primaryVersion.Uid)
}

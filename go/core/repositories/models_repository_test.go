package repositories

import (
    "testing";

    "github.com/marekgalovic/photon/go/core/storage";

    "github.com/stretchr/testify/suite";
)

type ModelsRepositoryTest struct {
    suite.Suite
    db *storage.Mysql
    repository *modelsRepository
    modelId int64
    versionId int64
    otherVersionId int64
    featureSetId int64
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
    result, err := test.db.Exec(
        `INSERT INTO models (name, runner_type) VALUES (?,?)`,
        "test_model", "pmml",
    )
    test.Require().Nil(err, "Failed to seed models.")

    test.modelId, err = result.LastInsertId()
    test.Require().Nil(err, "Failed to seed models.")

    result, err = test.db.Exec(`INSERT INTO model_versions (model_id, name, file_name, is_primary, is_shadow) VALUES (?,?,?,?,?)`, test.modelId, "version_1", "test.xml", true, false)
    test.Require().Nil(err, "Failed to seed model versions.")
    test.versionId, err = result.LastInsertId()
    test.Require().Nil(err, "Failed to seed model versions.")

    result, err = test.db.Exec(`INSERT INTO model_versions (model_id, name, file_name, is_primary, is_shadow) VALUES (?,?,?,?,?)`, test.modelId, "version_2", "test.xml", false, false)
    test.Require().Nil(err, "Failed to seed model versions.")
    test.otherVersionId, err = result.LastInsertId()
    test.Require().Nil(err, "Failed to seed model versions.")

    result, err = test.db.Exec(
        `INSERT INTO feature_sets (name, lookup_keys) VALUES (?,?)`,
        "precomputed_features", "resource_id,email",
    )
    test.Require().Nil(err, "Failed to seed feature sets.")

    test.featureSetId, err = result.LastInsertId()
    test.Require().Nil(err, "Failed to seed feature sets.")

    _, err = test.db.Exec(
        `INSERT INTO model_features (model_id, name, alias, required) VALUES (?,?,?,?),(?,?,?,?)`,
        test.modelId, "x1", "x1", true,
        test.modelId, "x2", "x2_alias", false,
    )
    test.Require().Nil(err, "Failed to seed model features.")

    _, err = test.db.Exec(
        `INSERT INTO model_precomputed_features (model_id, feature_set_id, name, alias, required) VALUES (?,?,?,?,?)`,
        test.modelId, test.featureSetId, "x1", "x1_precomputed", true,
    )
    test.Require().Nil(err, "Failed to seed model precomputed features.")
}

func (test *ModelsRepositoryTest) TestList() {
    models, err := test.repository.List()
    test.Require().Nil(err)

    test.Equal(1, len(models))
}

func (test *ModelsRepositoryTest) TestFind() {
    model, err := test.repository.Find(test.modelId)
    test.Require().Nil(err)

    test.Equal("test_model", model.Name)
    test.Equal(2, len(model.Features))
    test.Equal(1, len(model.PrecomputedFeatures))
}

func (test *ModelsRepositoryTest) TestFindByName() {
    model, err := test.repository.FindByName("test_model")
    test.Require().Nil(err)

    test.Equal(test.modelId, model.Id)
}

func (test *ModelsRepositoryTest) TestCreate() {
    storage.AssertCountChanged(test.db, "models", 1, func() {
        storage.AssertCountChanged(test.db, "model_features", 1, func() {
            storage.AssertCountChanged(test.db, "model_precomputed_features", 1, func() {
                model := &Model{
                    Name: "test_model_2",
                    RunnerType: "pmml",
                    Features: []*ModelFeature{
                        {Name: "x1", Required: true},
                    },
                    PrecomputedFeatures: map[int64][]*ModelFeature{
                        test.featureSetId: []*ModelFeature{
                            {Name: "x2", Alias: "x2_precomputed", Required: false},
                        },
                    },
                }
                _, err := test.repository.Create(model)
                test.Require().Nil(err)
            })
        })
    })
}

func (test *ModelsRepositoryTest) TestDelete() {
    storage.AssertCountChanged(test.db, "models", -1, func() {
        storage.AssertCountChanged(test.db, "model_features", -2, func() {
            storage.AssertCountChanged(test.db, "model_precomputed_features", -1, func() {
                storage.AssertCountChanged(test.db, "model_versions", -2, func() {
                    err := test.repository.Delete(test.modelId)
                    test.Nil(err)
                })
            })
        })
    })
}

func (test *ModelsRepositoryTest) TestListVersions() {
    versions, err := test.repository.ListVersions(test.modelId)
    test.Require().Nil(err)

    test.Equal(2, len(versions))
}

func (test *ModelsRepositoryTest) TestFindVersion() {
    version, err := test.repository.FindVersion(test.versionId)
    test.Require().Nil(err)

    test.Equal("version_1", version.Name)
    test.True(version.IsPrimary)
    test.False(version.IsShadow)
}

func (test *ModelsRepositoryTest) TestPrimaryVersion() {
    version, err := test.repository.PrimaryVersion(test.modelId)
    test.Require().Nil(err)

    test.Equal(test.versionId, version.Id)
}

func (test *ModelsRepositoryTest) TestCreateVersion() {
    storage.AssertCountChanged(test.db, "model_versions", 1, func() {
        version := &ModelVersion{
            ModelId: test.modelId,
            Name: "version_3",
            FileName: "test.xml",
            IsPrimary: false,
            IsShadow: false,
        }
        _, err := test.repository.CreateVersion(version)
        test.Require().Nil(err)
    })
}

func (test *ModelsRepositoryTest) TestCreatePrimaryVersion() {
    oldPrimaryVersion, err := test.repository.PrimaryVersion(test.modelId)
    test.Require().Nil(err)

    version := &ModelVersion{
        ModelId: test.modelId,
        Name: "version_4",
        FileName: "test.xml",
        IsPrimary: true,
        IsShadow: false,
    }

    createdPrimaryVersionId, err := test.repository.CreateVersion(version)
    test.Require().Nil(err)

    newPrimaryVersion, err := test.repository.PrimaryVersion(test.modelId)
    test.Require().Nil(err)

    test.Equal(createdPrimaryVersionId, newPrimaryVersion.Id)
    test.NotEqual(oldPrimaryVersion.Id, newPrimaryVersion.Id)
}

func (test *ModelsRepositoryTest) TestDeleteVersion() {
    storage.AssertCountChanged(test.db, "model_versions", -1, func() {
        err := test.repository.DeleteVersion(test.otherVersionId)
        test.Nil(err)
    })
}

func (test *ModelsRepositoryTest) TestDeleteVersionWhenVersionIsPrimary() {
    storage.AssertCountChanged(test.db, "model_versions", 0, func() {
        err := test.repository.DeleteVersion(test.versionId)
        test.NotNil(err)
    })
}

func (test *ModelsRepositoryTest) TestSetPrimaryVersion() {
    primaryVersion, err := test.repository.PrimaryVersion(test.modelId)
    test.Require().Nil(err)
    test.Equal(test.versionId, primaryVersion.Id)

    err = test.repository.SetPrimaryVersion(test.modelId, test.otherVersionId)
    test.Nil(err)

    primaryVersion, err = test.repository.PrimaryVersion(test.modelId)
    test.Require().Nil(err)
    test.Equal(test.otherVersionId, primaryVersion.Id)
}

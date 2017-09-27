package evaluator

import (
    "testing";

    "github.com/marekgalovic/photon/go/core/repositories";
    "github.com/marekgalovic/photon/go/core/repositories/mock";

    "github.com/stretchr/testify/suite";
    "github.com/golang/mock/gomock";
)

type ModelResolverTest struct {
    suite.Suite
    mockController *gomock.Controller
    modelsRepository *mock_repositories.MockModelsRepository
    resolver ModelResolver
    model *repositories.Model
    version *repositories.ModelVersion
}

func TestModelResolver(t *testing.T) {
    suite.Run(t, new(ModelResolverTest))
}

func (test *ModelResolverTest) SetupTest() {
    test.mockController = gomock.NewController(test.T())
    test.modelsRepository = mock_repositories.NewMockModelsRepository(test.mockController)
    test.resolver = NewModelResolver(test.modelsRepository)

    test.model = &repositories.Model{Id: 1, Name: "test"}
    test.version = &repositories.ModelVersion{Id: 2, ModelId: 1, Name: "test_version"}
}

func (test *ModelResolverTest) TearDownTest() {
    test.mockController.Finish()
}

func (test *ModelResolverTest) TestGetModel() {
    test.modelsRepository.EXPECT().FindByName("test").Return(test.model, nil)
    test.modelsRepository.EXPECT().PrimaryVersion(test.model.Id).Return(test.version, nil)

    model, version, err := test.resolver.GetModel("test")

    test.Require().Nil(err)
    test.Equal(test.model, model)
    test.Equal(test.version, version)
}

func (test *ModelResolverTest) TestGetModelCachesResults() {
    test.modelsRepository.EXPECT().FindByName("test").Return(test.model, nil)
    test.modelsRepository.EXPECT().PrimaryVersion(test.model.Id).Return(test.version, nil)

    model, version, err := test.resolver.GetModel("test")
    test.Require().Nil(err)
    test.Equal(test.model, model)
    test.Equal(test.version, version)

    model, version, err = test.resolver.GetModel("test")
    test.Require().Nil(err)
    test.Equal(test.model, model)
    test.Equal(test.version, version)
}

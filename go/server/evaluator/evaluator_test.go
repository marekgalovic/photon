package evaluator

import (
    "testing";

    "github.com/marekgalovic/photon/go/server/evaluator/mock";
    "github.com/marekgalovic/photon/go/core/repositories";
    "github.com/marekgalovic/photon/go/core/repositories/mock";
    "github.com/marekgalovic/photon/go/core/utils";
    pb "github.com/marekgalovic/photon/go/core/protos";
    mock_pb "github.com/marekgalovic/photon/go/core/protos/mock";

    "github.com/patrickmn/go-cache";
    "github.com/stretchr/testify/suite";
    "github.com/golang/mock/gomock";
)

type EvaluatorTest struct{
    suite.Suite
    mockController *gomock.Controller
    instancesRepository repositories.InstancesRepository
    modelResolver *mock_evaluator.MockModelResolver
    featuresResolver *mock_evaluator.MockFeaturesResolver
    runnerClient *mock_pb.MockRunnerServiceClient
    evaluator *evaluator
    model *repositories.Model
    modelVersion *repositories.ModelVersion
}

func TestEvaluator(t *testing.T) {
    suite.Run(t, new(EvaluatorTest))
}

func (test *EvaluatorTest) SetupSuite() {
    test.mockController = gomock.NewController(test.T())
    test.instancesRepository = mock_repositories.NewMockInstancesRepository(test.mockController)
    test.modelResolver = mock_evaluator.NewMockModelResolver(test.mockController)
    test.featuresResolver = mock_evaluator.NewMockFeaturesResolver(test.mockController)
    test.runnerClient = mock_pb.NewMockRunnerServiceClient(test.mockController)
    test.evaluator = NewEvaluator(test.modelResolver, test.featuresResolver, test.instancesRepository)
    // Stub runner client
    test.evaluator.clientsCache.Set("1", &clientsCacheEntry{client: test.runnerClient}, cache.NoExpiration)

    test.model = &repositories.Model{
        Id: 1,
        Name: "test",
        Features: []*repositories.ModelFeature{
            {Name: "x1", Alias: "x1", Required: true},
            {Name: "x2", Alias: "x2", Required: false},
        },
    }
    test.modelVersion = &repositories.ModelVersion{
        Id: 1,
        ModelId: 1,
        Name: "test_version",
        IsPrimary: true,
        IsShadow: false,
    }
}

func (test *EvaluatorTest) TearDownSuite() {
    test.mockController.Finish()
}

func (test *EvaluatorTest) TestEvaluate() {
    features := map[string]interface{}{"x1": 1, "x2": 3.4}
    result := map[string]interface{}{"y": 0.5}
    callFeatures, err := utils.InterfaceMapToValueInterfacePb(features)
    test.Require().Nil(err)
    callResult, err := utils.InterfaceMapToValueInterfacePb(result)
    test.Require().Nil(err)

    test.runnerClient.EXPECT().Evaluate(gomock.Any(), &pb.RunnerEvaluateRequest{VersionUid: "1", Features: callFeatures}).Return(&pb.RunnerEvaluateResponse{Result: callResult}, nil)
    test.modelResolver.EXPECT().GetModel("test").Return(test.model, test.modelVersion, nil)
    test.featuresResolver.EXPECT().Resolve(test.model, features).Return(features, nil)

    actualResult, err := test.evaluator.Evaluate("test", features)

    test.Require().Nil(err)
    test.Equal(actualResult, result)
}

package evaluator

import (
    "fmt";
    "time";
    "golang.org/x/net/context";

    "github.com/marekgalovic/photon/go/core/balancer";
    "github.com/marekgalovic/photon/go/core/repositories";
    "github.com/marekgalovic/photon/go/core/metrics";
    "github.com/marekgalovic/photon/go/core/utils";
    pb "github.com/marekgalovic/photon/go/core/protos";

    "google.golang.org/grpc";
    "github.com/patrickmn/go-cache";

    log "github.com/Sirupsen/logrus"
)

type Evaluator interface {
    Close()
    Evaluate(string, map[string]interface{}) (map[string]interface{}, error)
}

type evaluator struct {
    modelResolver ModelResolver
    featuresResolver FeaturesResolver
    instancesRepository repositories.InstancesRepository
    instancesResolver *balancer.Resolver
    clientsCache *cache.Cache
}

type clientsCacheEntry struct {
    conn *grpc.ClientConn
    client pb.RunnerServiceClient
}

func NewEvaluator(modelResolver ModelResolver, featuresResolver FeaturesResolver, instancesRepository repositories.InstancesRepository) *evaluator {
    return &evaluator{
        modelResolver: modelResolver,
        featuresResolver: featuresResolver,
        instancesRepository: instancesRepository,
        instancesResolver: balancer.NewResolver(instancesRepository),
        clientsCache: cache.New(1 * time.Minute, 1 * time.Minute),
    }
}

func (e *evaluator) Close() {
    for _, item := range e.clientsCache.Items() {
        entry := item.Object.(*clientsCacheEntry)
        entry.conn.Close()
    }
}

func (e *evaluator) Evaluate(modelName string, requestParams map[string]interface{}) (map[string]interface{}, error) {
    defer metrics.Runtime("evaluator.runtime", []string{fmt.Sprintf("model_name:%s", modelName), "method:evaluate"})()

    model, version, err := e.modelResolver.GetModel(modelName)
    if err != nil {
        return nil, err
    }

    features, err := e.featuresResolver.Resolve(model, requestParams)
    if err != nil {
        return nil, err
    }

    return e.call(version, features)
}

func (e *evaluator) call(version *repositories.ModelVersion, features map[string]interface{}) (map[string]interface{}, error) {
    defer metrics.Runtime("evaluator.runtime", []string{"method:call"})()

    client, err := e.getClient(version)
    if err != nil {
        return nil, err
    }

    featuresPb, err := utils.InterfaceMapToValueInterfacePb(features)
    if err != nil {
        return nil, err
    }

    log.Info(featuresPb)

    result, err := client.Evaluate(context.Background(), &pb.RunnerEvaluateRequest{
        VersionUid: fmt.Sprintf("%d", version.Id),
        Features: featuresPb,
    })
    if err != nil {
        return nil, err
    }

    return utils.ValueInterfacePbToInterfaceMap(result.Result)
}
 
func (e *evaluator) getClient(version *repositories.ModelVersion) (pb.RunnerServiceClient, error) {
    if cached, exists := e.clientsCache.Get(version.StringId()); exists {
        entry := cached.(*clientsCacheEntry)
        return entry.client, nil
    }

    conn, err := grpc.Dial(version.StringId(), grpc.WithInsecure(), grpc.WithBalancer(grpc.RoundRobin(e.instancesResolver)))
    if err != nil {
        return nil, fmt.Errorf("Failed to create instance connection. %v", err)
    }

    client := pb.NewRunnerServiceClient(conn)
    e.clientsCache.Set(version.StringId(), &clientsCacheEntry{conn: conn, client: client}, cache.DefaultExpiration)

    return client, nil
}

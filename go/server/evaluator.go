package server

import (
    "fmt";
    "time";
    "golang.org/x/net/context";

    "github.com/marekgalovic/photon/go/core/balancer";
    "github.com/marekgalovic/photon/go/core/storage/repositories";
    "github.com/marekgalovic/photon/go/core/metrics";
    pb "github.com/marekgalovic/photon/go/core/protos";

    "google.golang.org/grpc";
    "github.com/patrickmn/go-cache";
    log "github.com/Sirupsen/logrus"
)

type Evaluator struct {
    modelResolver *ModelResolver
    featuresResolver *FeaturesResolver
    instancesRepository *repositories.InstancesRepository
    instancesResolver *balancer.Resolver
    clientsCache *cache.Cache
}

type clientsCacheEntry struct {
    conn *grpc.ClientConn
    client pb.RunnerServiceClient
}

func NewEvaluator(modelResolver *ModelResolver, featuresResolver *FeaturesResolver, instancesRepository *repositories.InstancesRepository) *Evaluator {
    return &Evaluator{
        modelResolver: modelResolver,
        featuresResolver: featuresResolver,
        instancesRepository: instancesRepository,
        instancesResolver: balancer.NewResolver(instancesRepository),
        clientsCache: cache.New(1 * time.Minute, 1 * time.Minute),
    }
}

func (e *Evaluator) Close() {
    for _, item := range e.clientsCache.Items() {
        entry := item.Object.(*clientsCacheEntry)
        entry.conn.Close()
    }
}

func (e *Evaluator) Evaluate(modelName string, requestParams map[string]interface{}) (map[string]interface{}, error) {
    defer metrics.Runtime("evaluator.runtime", []string{fmt.Sprintf("model_name:%s", modelName), "method:evaluate"})()

    _, version, err := e.modelResolver.GetModel(modelName)
    if err != nil {
        return nil, err
    }

    features, err := e.featuresResolver.Resolve(version, requestParams)
    if err != nil {
        return nil, err
    }

    return e.call(version.Uid, features)
}

func (e *Evaluator) call(versionUid string, features map[string]interface{}) (map[string]interface{}, error) {
    defer metrics.Runtime("evaluator.runtime", []string{fmt.Sprintf("model_version_uid:%s", versionUid), "method:call"})()

    client, err := e.getClient(versionUid)
    if err != nil {
        return nil, err
    }

    result, err := client.Evaluate(context.Background(), &pb.RunnerEvaluateRequest{})
    if err != nil {
        return nil, err
    }

    log.Info("Runner result: ", result)
    return nil, nil
}
 
func (e *Evaluator) getClient(versionUid string) (pb.RunnerServiceClient, error) {
    if cached, exists := e.clientsCache.Get(versionUid); exists {
        entry := cached.(*clientsCacheEntry)
        return entry.client, nil
    }

    conn, err := grpc.Dial(versionUid, grpc.WithInsecure(), grpc.WithBalancer(grpc.RoundRobin(e.instancesResolver)))
    if err != nil {
        return nil, fmt.Errorf("Failed to create instance connection. %v", err)
    }

    client := pb.NewRunnerServiceClient(conn)
    e.clientsCache.Set(versionUid, &clientsCacheEntry{conn: conn, client: client}, cache.DefaultExpiration)

    return client, nil
}

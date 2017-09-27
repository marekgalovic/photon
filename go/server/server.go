package server

import (
    "fmt";
    "net";

    "github.com/marekgalovic/photon/go/core/storage";
    "github.com/marekgalovic/photon/go/core/storage/features";
    "github.com/marekgalovic/photon/go/core/storage/files";
    "github.com/marekgalovic/photon/go/core/cluster";
    "github.com/marekgalovic/photon/go/core/repositories";
    pb "github.com/marekgalovic/photon/go/core/protos";
    "github.com/marekgalovic/photon/go/server/evaluator";
    "github.com/marekgalovic/photon/go/server/services";

    "google.golang.org/grpc";
    "google.golang.org/grpc/credentials";
)

type Server interface {
    Serve()
    Stop()
}

type ServerConfig struct {
    Address string
    Port int
    Tls credentials.TransportCredentials
}

type server struct {
    config ServerConfig
    mysql *storage.Mysql
    zookeeper *storage.Zookeeper
    kubernetes cluster.Kubernetes
    featuresStore features.FeaturesStore
    filesStore files.FilesStore
    credentialsRepository repositories.CredentialsRepository
    featuresRepository repositories.FeaturesRepository
    modelsRepository repositories.ModelsRepository
    deployerRepository repositories.DeployerRepository
    instancesRepository repositories.InstancesRepository
    modelResolver evaluator.ModelResolver
    featuresResolver evaluator.FeaturesResolver
    evaluator evaluator.Evaluator
    listener net.Listener
    grpcServer *grpc.Server
}

func NewServer(config ServerConfig, mysql *storage.Mysql, zookeeper *storage.Zookeeper, kubernetes cluster.Kubernetes, featuresStore features.FeaturesStore, filesStore files.FilesStore) *server {
    s := &server {
        config: config,
        mysql: mysql,
        zookeeper: zookeeper,
        kubernetes: kubernetes,
        featuresStore: featuresStore,
        filesStore: filesStore,
    }

    s.initializeGrpcServer()
    s.initializeRepositories()
    s.initializeEvaluator()
    s.registerServices()

    return s
}

func (s *server) Serve() error {
    var err error

    s.listener, err = net.Listen("tcp", s.bindAddress())
    if err != nil {
        return err
    }

    go s.grpcServer.Serve(s.listener)
    return nil
}

func (s *server) Stop() {
    s.evaluator.Close()
    s.grpcServer.GracefulStop()
    s.listener.Close()
}

func (s *server) initializeGrpcServer() {
    grpcServerOptions := make([]grpc.ServerOption, 0)
    if s.config.Tls != nil {
        grpcServerOptions = append(grpcServerOptions, grpc.Creds(s.config.Tls))
    }

    s.grpcServer = grpc.NewServer(grpcServerOptions...)
}

func (s *server) initializeRepositories() {
    s.credentialsRepository = repositories.NewCredentialsRepository(s.mysql)
    s.featuresRepository = repositories.NewFeaturesRepository(s.mysql)
    s.modelsRepository = repositories.NewModelsRepository(s.mysql)
    s.deployerRepository = repositories.NewDeployerRepository(s.zookeeper)
    s.instancesRepository = repositories.NewInstancesRepository(s.zookeeper)
}

func (s *server) initializeEvaluator() {
    s.modelResolver = evaluator.NewModelResolver(s.modelsRepository)
    s.featuresResolver = evaluator.NewFeaturesResolver(s.featuresRepository, s.featuresStore)
    s.evaluator = evaluator.NewEvaluator(s.modelResolver, s.featuresResolver, s.instancesRepository)
}

func (s *server) registerServices() {
    pb.RegisterEvaluatorServiceServer(s.grpcServer, services.NewEvaluatorService(s.evaluator))
    pb.RegisterModelsServiceServer(s.grpcServer, services.NewModelsService(s.modelsRepository, s.deployerRepository, s.filesStore, s.kubernetes))
    pb.RegisterFeaturesServiceServer(s.grpcServer, services.NewFeaturesService(s.featuresRepository))
}

func (s *server) bindAddress() string {
    return net.JoinHostPort(s.config.Address, fmt.Sprintf("%d", s.config.Port))
}

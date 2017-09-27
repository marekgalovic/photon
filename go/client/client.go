package photon

import (
    "google.golang.org/grpc";
)

type Client interface {
    ModelsService
    EvaluatorService
}

type client struct {
    config *Config
    conn *grpc.ClientConn
    *modelsService
    *evaluatorService
}

func NewClient(config *Config, credentials *credentials) (*client, error) {
    conn, err := grpc.Dial(config.serverAddr(), grpc.WithPerRPCCredentials(credentials), grpc.WithInsecure())
    if err != nil {
        return nil, err
    }

    return &client {
        config: config,
        conn: conn,
        evaluatorService: newEvaluatorService(conn),
        modelsService: newModelsService(conn),
    }, nil
}

func (c *client) Close() {
    c.conn.Close()
}

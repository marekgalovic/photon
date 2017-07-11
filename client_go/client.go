package serving

import (
    "golang.org/x/net/context";
    "google.golang.org/grpc";

    pb "github.com/marekgalovic/serving/client_go/protos"
)

type Client struct {
    config *Config
    conn *grpc.ClientConn
    evaluatorClient pb.EvaluatorServiceClient
}

func NewClient(config *Config) (*Client, error) {
    conn, err := grpc.Dial(config.serverAddr(), grpc.WithInsecure())
    if err != nil {
        return nil, err
    }

    return &Client {
        config: config,
        conn: conn,
        evaluatorClient: pb.NewEvaluatorServiceClient(conn),
    }, nil
}

func (c *Client) Close() {
    c.conn.Close()
}

func (c *Client) Evaluate(modelUid string, features map[string]interface{}) (map[string]interface{}, error) {
    valueInterfaces, err := interfaceMapToValueInterfacePb(features)
    if err != nil {
        return nil, err
    }

    result, err := c.evaluatorClient.Evaluate(context.Background(), &pb.EvaluationRequest{modelUid, valueInterfaces})
    if err != nil {
        return nil, err
    }

    return valueInterfacePbToInterfaceMap(result.Result)
}

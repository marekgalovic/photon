package photon

import (
    "golang.org/x/net/context";
    // "google.golang.org/grpc";
)

type credentials struct {
    key string
    secret string
}

func NewCredentials(key, secret string) *credentials {
    return &credentials{key: key, secret: secret}
}

func (c *credentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
    return map[string]string{"key": c.key, "secret": c.secret}, nil
}

func (c *credentials) RequireTransportSecurity() bool {
    return false
}

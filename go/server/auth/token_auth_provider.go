package auth

import (
    "fmt";
    "golang.org/x/net/context";

    "github.com/marekgalovic/photon/go/core/repositories";

    "google.golang.org/grpc/metadata";
)

type tokenAuthProvider struct {
    credentialsRepository repositories.CredentialsRepository
}

func NewTokenAuthProvider(credentialsRepository repositories.CredentialsRepository) *tokenAuthProvider {
    return &tokenAuthProvider{
        credentialsRepository: credentialsRepository,
    }
}

func (p *tokenAuthProvider) Authenticate(ctx context.Context) (context.Context, error) {
    key, secret, err := p.getCredentialsFromContext(ctx)
    if err != nil {
        return ctx, err
    }

    credential, err := p.credentialsRepository.Find(key)
    if err != nil {
        return ctx, err
    }

    if err := credential.Verify(secret); err != nil {
        return ctx, fmt.Errorf("Unauthorized.")
    }
    
    return ctx, nil
}

func (p *tokenAuthProvider) getCredentialsFromContext(ctx context.Context) (string, string, error) {
    md, ok := metadata.FromIncomingContext(ctx)
    if !ok {
        return "", "", fmt.Errorf("Failed to get request metadata.")
    }
    key, exists := md["key"]
    if !exists || (len(key) != 1) || key[0] == "" {
        return "", "", fmt.Errorf("Missing or invalid key.")
    }
    secret, exists := md["secret"]
    if !exists || (len(secret) != 1) || secret[0] == "" {
        return "", "", fmt.Errorf("Missing or invalid secret.")
    }

    return key[0], secret[0], nil
}

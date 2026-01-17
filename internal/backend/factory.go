package backend

import (
	"context"

	"github.com/hashicorp/vault/sdk/logical"
)

func InternalFactory(ctx context.Context, conf *logical.BackendConfig) (logical.Backend, error) {
	b := newSolanaBackend()
	if err := b.Setup(ctx, conf); err != nil {
		return nil, err
	}
	return b, nil
}

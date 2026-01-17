package backend

import (
	"context"
	"strings"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

const (
	backendHelp = `
The Solana secrets backend allows for the dynamic and secure
creation, access, and management of wallet keypair material.
`
)

type SolanaBackend struct {
	*framework.Backend
}

type WalletEntry struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

func newSolanaBackend() *SolanaBackend {
	var s = SolanaBackend{}
	s.Backend = &framework.Backend{
		Help: strings.TrimSpace(backendHelp),
		PathsSpecial: &logical.Paths{
			LocalStorage: []string{},
			SealWrapStorage: []string{
				"wallet/",
			},
		},
		Paths: framework.PathAppend(
			pathWallet(&s),
		),
		Secrets:     []*framework.Secret{},
		BackendType: logical.TypeLogical,
	}
	return &s
}

func (s *SolanaBackend) getWallet(ctx context.Context, store logical.Storage, id string) (*WalletEntry, error) {
	entry, err := store.Get(ctx, "wallet/"+id)
	if err != nil {
		return nil, err
	}

	if entry == nil {
		return nil, nil
	}

	var wallet WalletEntry
	if err := entry.DecodeJSON(&wallet); err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (s *SolanaBackend) setWallet(ctx context.Context, store logical.Storage, id string, w *WalletEntry) error {
	entry, err := logical.StorageEntryJSON("wallet/"+id, w)
	if err != nil {
		return err
	}

	return store.Put(ctx, entry)
}

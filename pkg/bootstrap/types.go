package bootstrap

import (
	"context"
	"errors"

	"github.com/open-panoptes/opni/pkg/ident"
	"github.com/open-panoptes/opni/pkg/keyring"
	"github.com/open-panoptes/opni/pkg/storage"
)

type Bootstrapper interface {
	Bootstrap(context.Context, ident.Provider) (keyring.Keyring, error)
}

type Storage interface {
	storage.TokenStore
	storage.ClusterStore
	storage.KeyringStoreBroker
}

type StorageConfig struct {
	storage.TokenStore
	storage.ClusterStore
	storage.KeyringStoreBroker
}

var (
	ErrInvalidEndpoint    = errors.New("invalid endpoint")
	ErrNoRootCA           = errors.New("no root CA found in peer certificates")
	ErrLeafNotSigned      = errors.New("leaf certificate not signed by the root CA")
	ErrKeyExpired         = errors.New("key expired")
	ErrRootCAHashMismatch = errors.New("root CA hash mismatch")
	ErrNoValidSignature   = errors.New("no valid signature found in response")
	ErrNoToken            = errors.New("no bootstrap token provided")
)

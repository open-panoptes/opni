package local

import (
	"context"

	"github.com/rancher/opni/pkg/storage"
	"github.com/rancher/opni/pkg/storage/kvutil"
	"github.com/rancher/opni/pkg/util"
	"golang.org/x/crypto/bcrypt"
)

const (
	difficulty     = 10
	passwordLength = 20
	storageKey     = "local_admin"
)

type LocalAuthenticator interface {
	GenerateAdminPassword(context.Context) (string, error)
	ComparePassword(ctx context.Context, pw []byte) error
}

func NewLocalAuthenticator(kv storage.KeyValueStore) LocalAuthenticator {
	return &implLocalAuthenticator{
		store: kvutil.WithKey[[]byte](kv, storageKey),
	}
}

type implLocalAuthenticator struct {
	store storage.ValueStoreT[[]byte]
}

func (a *implLocalAuthenticator) GenerateAdminPassword(ctx context.Context) (string, error) {
	pw := util.GenerateRandomString(passwordLength)
	hash, err := bcrypt.GenerateFromPassword(pw, difficulty)
	if err != nil {
		return "", err
	}
	err = a.store.Put(ctx, hash)
	if err != nil {
		return "", err
	}
	return string(pw), nil
}

func (a *implLocalAuthenticator) ComparePassword(ctx context.Context, pw []byte) error {
	hash, err := a.store.Get(ctx)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword(hash, pw)
}

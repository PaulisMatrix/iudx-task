package storage

import (
	"context"
)

type StoreIface interface {
	Set(ctx context.Context, data []byte) error
	Get(ctx context.Context) ([]byte, error)
	Delete(ctx context.Context) error
}

type ClientStore struct {
	store StoreIface
}

var _ StoreIface = (*ClientStore)(nil)

func NewClientStore(store StoreIface) *ClientStore {
	return &ClientStore{
		store: store,
	}
}

func (c *ClientStore) Set(ctx context.Context, data []byte) error {
	return nil
}

func (c *ClientStore) Get(ctx context.Context) ([]byte, error) {
	return []byte{}, nil
}

func (c *ClientStore) Delete(ctx context.Context) error {
	return nil
}

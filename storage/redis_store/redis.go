package redis_store

import (
	"context"
	store "datakaveri/storage"

	"github.com/gomodule/redigo/redis"
)

type RedisStore struct {
	redisPool *redis.Pool
}

var _ store.StoreIface = (*RedisStore)(nil)

func NewRedisStore() *RedisStore {
	return &RedisStore{
		redisPool: NewRedisPool("localhost:6379", ""),
	}
}

func (r *RedisStore) Set(ctx context.Context, data []byte) error {
	return nil
}

func (r *RedisStore) Get(ctx context.Context) ([]byte, error) {
	return []byte{}, nil
}

func (r *RedisStore) Delete(ctx context.Context) error {
	return nil
}

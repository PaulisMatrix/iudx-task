package redis_store

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

func NewRedisPool(host, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				host,
				redis.DialDatabase(0),
				redis.DialPassword(password),
				redis.DialConnectTimeout(5*time.Second),
				redis.DialReadTimeout(10*time.Second),
				redis.DialWriteTimeout(10*time.Second),
			)
		},
	}
}

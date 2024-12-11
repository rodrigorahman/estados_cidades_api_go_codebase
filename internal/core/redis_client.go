package core

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisClientInterface interface {
	Set(key string, value interface{}, expiration int) error
	Get(key string) (string, error)
	Ping() error
	Close() error
}

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisClient(addr, password string, db int) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisClient{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (r *RedisClient) Ping() error {
	_, err := r.client.Ping(r.ctx).Result()
	return err
}

func (r *RedisClient) Set(key string, value interface{}, ttl time.Duration) error {
	return r.client.Set(r.ctx, key, value, ttl).Err()
}

func (r *RedisClient) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}
func (r *RedisClient) Close() error {
	return r.client.Close()
}

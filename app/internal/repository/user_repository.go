package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(addr string, password string, db int) *RedisRepository {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisRepository{
		client: client,
	}
}

func (r *RedisRepository) Get(
	ctx context.Context,
	key string,
) (string, error) {

	return r.client.Get(ctx, key).Result()
}

func (r *RedisRepository) Set(
	ctx context.Context,
	key string,
	value string,
	expiration time.Duration,
) error {

	return r.client.Set(
		ctx,
		key,
		value,
		expiration,
	).Err()
}

func (r *RedisRepository) Delete(
	ctx context.Context,
	key string,
) error {

	return r.client.Del(ctx, key).Err()
}

func (r *RedisRepository) Ping(
	ctx context.Context,
) error {

	return r.client.Ping(ctx).Err()
}
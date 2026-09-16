package cache

import (
	"context"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/os_config"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache() *RedisCache {
	conf := os_config.NewEnvironmentConfig()
	db, err := strconv.Atoi(conf.RedisDB)
	if err != nil {
		panic("invalid config for  redisdb")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddress,
		Password: conf.RedisPassword,
		DB:       db,
	})

	return &RedisCache{
		client: client,
	}
}

func (r *RedisCache) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

func (r *RedisCache) Set(
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

func (r *RedisCache) Get(
	ctx context.Context,
	key string,
) (string, error) {

	return r.client.Get(ctx, key).Result()
}

func (r *RedisCache) Delete(
	ctx context.Context,
	key string,
) error {

	return r.client.Del(ctx, key).Err()
}

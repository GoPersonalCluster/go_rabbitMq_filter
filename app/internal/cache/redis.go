package cache

import (
	"context"
	"strconv"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/config"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

func NewRedis() (*Redis, error) {
	conf := config.NewEnvironmentConfig()
	db, err := strconv.Atoi(conf.RedisDB)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddress,
		Password: conf.RedisPassword,
		DB:       db,
	})

	return &Redis{
		Client: client,
	}, nil
}

func (r *Redis) Ping(ctx context.Context) error {
	return r.Client.Ping(ctx).Err()
}

// func RedisHandler(cacheKey string, rc *redis.Client) {
// 	ctx := context.Background()

// 	err = rc.Set(
// 		ctx,
// 		cacheKey,
// 		data,
// 		5*time.Minute,
// 	).Err()

// }

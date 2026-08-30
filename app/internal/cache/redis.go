package cache

import (
	"context"
	"os"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/config"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

func NewRedis() *Redis {
	conf := config.NewEnvironmentConfig()


	client := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddress,
		Password: conf.RedisPassword,
		DB:       conf.RedisDB,
	})

	return &Redis{
		Client: client,
	}
}

func (r *Redis) Ping(ctx context.Context) error {
	return r.Client.Ping(ctx).Err()
}

func RedisHandler(cacheKey string ,rc *redis.Client ){
	ctx := context.Background()

		err = rc.Set(
		ctx,
		cacheKey,
		data,
		5*time.Minute,
	).Err()


}

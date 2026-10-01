package cache

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/cache"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestRedisCache(t *testing.T) (*cache.RedisCache, *miniredis.Miniredis) {
	t.Helper()

	server := miniredis.RunT(t)

	cache := cache.NewRedisCache()

	return cache, server
}

func TestRedisCache_Ping(t *testing.T) {
	cache, _ := newTestRedisCache(t)

	err := cache.Ping(context.Background())

	require.NoError(t, err)
}

type testStructure struct {
	value int `json:"value"`
}

func TestRedisCache_Set(t *testing.T) {
	cache, server := newTestRedisCache(t)

	ctx := context.Background()

	testStruct := testStructure{
		value: 1,
	}
	data, err := json.Marshal(testStruct)
	text := string(data)

	cache.Set(
		ctx,
		"user:1",
		text,
		0,
	)

	require.NoError(t, err)

	value, err := cache.Get(ctx, "user:1")
	println("======================================saved cache")
	println(value)

	require.NoError(t, err)
	assert.NotNil(t, server)
}

func TestRedisCache_Get(t *testing.T) {
	cache, server := newTestRedisCache(t)

	ctx := context.Background()

	server.Set("user:1", "Walter")

	value, err := cache.Get(
		ctx,
		"user:1",
	)

	println("======================================saved cache")
	println(value)

	require.NoError(t, err)
	assert.Equal(t, "Walter", value)
}

func TestRedisCache_Get_KeyNotFound(t *testing.T) {
	cache, _ := newTestRedisCache(t)

	ctx := context.Background()

	_, err := cache.Get(
		ctx,
		"user:999",
	)

	require.Error(t, err)
	assert.ErrorIs(t, err, redis.Nil)
}

func TestRedisCache_Delete(t *testing.T) {
	cache, server := newTestRedisCache(t)

	ctx := context.Background()

	server.Set("user:1", "Walter")

	err := cache.Delete(
		ctx,
		"user:1",
	)

	require.NoError(t, err)

	_, err = cache.Get(
		ctx,
		"user:1",
	)

	assert.ErrorIs(t, err, redis.Nil)
}

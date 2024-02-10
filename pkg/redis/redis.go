//go:generate mockgen -destination ../../mocks/mock_redis.go -package mocks github.com/saufiroja/cqrs/pkg/redis IRedis
package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/saufiroja/cqrs/config"
	"github.com/saufiroja/cqrs/pkg/logger"
	"time"
)

type IRedis interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) error
	Del(key string) error
}

type Redis struct {
	*redis.Client
}

func NewRedis(conf *config.AppConfig, log *logger.Logger) IRedis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Redis.Host, conf.Redis.Port),
		Password: "",
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.StartLogger("redis.go", "NewRedis").Error("error connecting to redis")
		panic(err)
	}

	log.StartLogger("redis.go", "NewRedis").Info("connected to redis")

	return &Redis{rdb}
}

func (r *Redis) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.Client.Get(ctx, key).Result()
}

func (r *Redis) Set(key string, value interface{}, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.Client.Set(ctx, key, value, expiration).Err()
}

func (r *Redis) Del(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.Client.Del(ctx, key).Err()
}

func (r *Redis) Close(ctx context.Context) {
	err := r.Client.Close()
	if err != nil {
		panic(err)
	}
}

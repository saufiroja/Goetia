package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/saufiroja/cqrs/config"
	"github.com/saufiroja/cqrs/pkg/logger"
	"time"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(conf *config.AppConfig, log *logger.Logger) *Redis {
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

	return &Redis{client: rdb}
}

func (r *Redis) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.client.Get(ctx, key).Result()
}

func (r *Redis) Set(key string, value interface{}, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *Redis) Del(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.client.Del(ctx, key).Err()
}

func (r *Redis) Close(ctx context.Context) {
	err := r.client.Close()
	if err != nil {
		panic(err)
	}
}

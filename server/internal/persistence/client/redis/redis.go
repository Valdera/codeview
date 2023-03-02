package client

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Configuration struct {
	Host     string
	Port     string
	Password string
}

func (cfg *Configuration) GetAddress() string {
	return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
}

type RedisDriver interface {
	Get(context context.Context, key string) *redis.StringCmd
	Set(context context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Del(context context.Context, keys ...string) *redis.IntCmd
	Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd
}

func Init(cfg Configuration) (RedisDriver, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.GetAddress(),
		Password: cfg.Password,
	})
	return client, nil
}

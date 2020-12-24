package database

import (
	. "GoEchoton/config"
	"context"

	"github.com/go-redis/redis/v8"
)

// Redis redis结构体
type Redis struct {
	client *redis.Client
}

var ctx context.Context = context.Background()

// Get Get命令
func (r *Redis) Get(key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

// NewRedis 创建Redis
func NewRedis() *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     Config.Redis.Host,
		Password: Config.Redis.Passwd,
		DB:       Config.Redis.DB,
	})
	return &Redis{client: rdb}
}

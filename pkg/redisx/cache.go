package redisx

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/sync/singleflight"
)

type Cache struct {
	client *redis.Client

	ttl time.Duration // 缓存过期时间

	sg singleflight.Group
}

// NewCache use *redis.Options created.
func NewCache(options *redis.Options, ttl time.Duration) *Cache {
	return &Cache{client: redis.NewClient(options), ttl: ttl}
}

// Get cache get. if nil, dofunc() set.
func (c *Cache) Get(ctx context.Context, key string, i interface{}, dofunc func(ctx context.Context) (interface{}, error)) error {
	// fast. 读取缓存
	v, err := c.client.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return err
	}
	if err == nil {
		err = json.Unmarshal([]byte(v), i)
		if err == nil {
			return nil
		}

		// 错误数据删除
		c.client.Del(ctx, key)
	}

	return c.getSlow(ctx, key, i, dofunc)
}

// getSlow 无缓存，需要dofunc创建
func (c *Cache) getSlow(ctx context.Context, key string, i interface{}, dofunc func(ctx context.Context) (interface{}, error)) error {
	// 请求合并，防止穿透
	_, err, _ := c.sg.Do(key, func() (interface{}, error) {
		ri, err := dofunc(ctx)
		if err != nil {
			return nil, err
		}

		b, err := json.Marshal(ri)
		if err != nil {
			return nil, err
		}
		c.client.Set(ctx, key, string(b), c.ttl)

		err = json.Unmarshal(b, i)
		return nil, err
	})

	return err
}

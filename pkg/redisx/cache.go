package redisx

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/sync/singleflight"
)

type Cache struct {
	client *redis.Client

	ttl time.Duration // 缓存过期时间

	sg singleflight.Group
}

type DoFunc func(ctx context.Context) (interface{}, error)

// NewCache use *redis.Options created.
func NewCache(options *redis.Options, ttl time.Duration) *Cache {
	return &Cache{client: redis.NewClient(options), ttl: ttl}
}

// Get cache get. if nil, dofunc() set.
func (c *Cache) Get(ctx context.Context, key string, i interface{}, dofunc DoFunc) error {
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
func (c *Cache) getSlow(ctx context.Context, key string, i interface{}, dofunc DoFunc) error {
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

func (c *Cache) HGetAll(ctx context.Context, key string, i interface{}, dofunc DoFunc) error {
	// fast. 读取缓存
	mapCmd := c.client.HGetAll(ctx, key)
	err := mapCmd.Err()
	if err == nil {
		if err := mapCmd.Scan(i); err != nil {
			return err
		}
		return nil
	} else if err != redis.Nil {
		return err
	}

	// slow 加载缓存
	_, err, _ = c.sg.Do(key, func() (interface{}, error) {
		ri, err := dofunc(ctx)
		if err != nil {
			return nil, err
		}

		// 反射获取
		rd := reflect.ValueOf(ri)
		if rd.Kind() == reflect.Ptr {
			i = ri
			rd = rd.Elem()
		} else {
			i = &ri
		}
		td := rd.Type()

		_, err = c.client.Pipelined(ctx, func(rdb redis.Pipeliner) error {
			for j := 0; j < rd.NumField(); j++ {
				v := rd.Field(j)
				t := td.Field(j)

				switch v.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					rdb.HSet(ctx, key, t.Tag.Get("redis"), v.Int())
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					rdb.HSet(ctx, key, t.Tag.Get("redis"), v.Uint())
				case reflect.Float32, reflect.Float64:
					rdb.HSet(ctx, key, t.Tag.Get("redis"), v.Float())
				case reflect.Bool:
					rdb.HSet(ctx, key, t.Tag.Get("redis"), v.Bool())
				case reflect.String:
					rdb.HSet(ctx, key, t.Tag.Get("redis"), v.String())
				default:
					return errors.New("no support type")
				}
			}
			return nil
		})
		return nil, err
	})
	return err
}

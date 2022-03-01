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

	ttl time.Duration // 缓存过期时间, 单位秒

	sg singleflight.Group
}

type DoFunc func(ctx context.Context) (interface{}, error)

// NewCache use *redis.Options created.
func NewCache(options *redis.Options, ttl int64) *Cache {
	return &Cache{client: redis.NewClient(options), ttl: time.Duration(ttl) * time.Second}
}

// Get 获取缓存，使用的redis.String类型缓存，编码解码使用json
func (c *Cache) Get(ctx context.Context, key string, i interface{}, dofunc DoFunc) error {
	id := reflect.ValueOf(i)
	if id.Kind() != reflect.Ptr {
		return errors.New("i is not a pointer")
	}

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

	return c.getSlow(ctx, key, id, dofunc)
}

// getSlow 无缓存，需要dofunc创建
func (c *Cache) getSlow(ctx context.Context, key string, id reflect.Value, dofunc DoFunc) error {
	// 请求合并，防止穿透
	r, err, _ := c.sg.Do(key, func() (interface{}, error) {
		ri, err := dofunc(ctx)
		if err != nil {
			return nil, err
		}

		b, err := json.Marshal(ri)
		if err != nil {
			return nil, err
		}
		c.client.Set(ctx, key, string(b), c.ttl)

		return ri, err
	})

	rd := reflect.ValueOf(r)
	if rd.Kind() == reflect.Ptr {
		rd = rd.Elem()
	}
	id.Elem().Set(rd)

	return err
}

// HGetAll 获取缓存，使用redis.Hash类型缓存，仅支持使用基础类型元素的Struct及其指针
func (c *Cache) HGetAll(ctx context.Context, key string, i interface{}, dofunc DoFunc) error {
	id := reflect.ValueOf(i)
	if id.Kind() != reflect.Ptr {
		return errors.New("i is not a pointer")
	}

	// fast. 读取缓存
	mapCmd := c.client.HGetAll(ctx, key)
	err := mapCmd.Err()
	if err == nil {
		if len(mapCmd.Val()) > 0 {
			if err := mapCmd.Scan(i); err != nil {
				return err
			}
			return nil
		}

	} else {
		return err
	}

	// slow 加载缓存
	r, err, _ := c.sg.Do(key, func() (interface{}, error) {
		ri, err := dofunc(ctx)
		if err != nil {
			return nil, err
		}

		// 反射获取
		rd := reflect.ValueOf(ri)
		if rd.Kind() == reflect.Ptr {
			rd = rd.Elem()
		}

		td := rd.Type()

		_, err = c.client.Pipelined(ctx, func(rdb redis.Pipeliner) error {
			if err := reflectStructHash(ctx, key, rdb, rd, td); err != nil {
				return err
			}
			rdb.Expire(ctx, key, c.ttl)
			return nil
		})
		return ri, err
	})

	rdv := reflect.ValueOf(r)
	if rdv.Kind() == reflect.Ptr {
		rdv = rdv.Elem()
	}
	id.Elem().Set(rdv)
	return err
}

// reflectStructHash 反射结构字段到redis.Hash
func reflectStructHash(ctx context.Context, key string, rdb redis.Pipeliner, rd reflect.Value, td reflect.Type) error {
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
		case reflect.Struct:
			if err := reflectStructHash(ctx, key, rdb, v, v.Type()); err != nil {
				return err
			}
		default:
			return errors.New("no support type")
		}

	}
	return nil
}

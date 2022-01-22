package redisx

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"gotest.tools/assert"
)

func TestCacheGet(t *testing.T) {
	cache := NewCache(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	}, 30*time.Second)

	var s struct {
		A string `json:"a"`
	}
	if err := cache.Get(context.Background(), "hello", &s, func(ctx context.Context) (interface{}, error) {
		var s1 struct {
			A string `json:"a"`
		}
		s1.A = "ppp"
		return s1, nil
	}); err != nil {
		t.Error(err)
	}

	assert.Equal(t, s.A, "ppp")

	for i := 0; i < 10; i++ {
		go func() {
			var ss string
			if err := cache.Get(context.Background(), "g1", &ss, func(ctx context.Context) (interface{}, error) {
				return "dafadad", nil
			}); err != nil {
				t.Error(err)
			}
		}()
	}
}

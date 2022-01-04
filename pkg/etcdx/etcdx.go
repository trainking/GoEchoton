package etcdx

import (
	"context"
	"errors"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var ErrorNotFound = errors.New("not found")

type ClientX struct {
	client *clientv3.Client
}

func New(etcdGateway []string) *ClientX {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdGateway,
		DialTimeout: 10 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	return &ClientX{client: client}
}

// Put 设置一个键值
func (c *ClientX) Put(ctx context.Context, key, value string) error {
	kv := clientv3.NewKV(c.client)
	_, err := kv.Put(ctx, key, value)
	return err
}

// Get 获取一个键值，找不到 ErrorNotFound
func (c *ClientX) Get(ctx context.Context, key string) (string, error) {
	kv := clientv3.NewKV(c.client)
	resp, err := kv.Get(ctx, key)
	if err != nil {
		return "", err
	}
	for _, _kv := range resp.Kvs {
		return string(_kv.Value), nil
	}
	return "", ErrorNotFound
}

// GetWithPrefix 前缀获取一组键值，找不到 ErrorNotFound
func (c *ClientX) GetWithPrefix(ctx context.Context, key string) ([]string, error) {
	kv := clientv3.NewKV(c.client)
	resp, err := kv.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	length := len(resp.Kvs)
	if length == 0 {
		return nil, ErrorNotFound
	}

	rList := make([]string, length)
	j := 0
	for _, v := range resp.Kvs {
		rList[j] = string(v.Value)
		j++
	}
	return rList, nil
}

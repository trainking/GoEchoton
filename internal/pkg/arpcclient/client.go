package arpcclient

import (
	"net"
	"time"

	"github.com/lesismal/arpc"
)

type Client struct {
	pool *arpc.ClientPool
}

// New 创建连接池
func New(listenAddr string, size int) (*Client, error) {
	pool, err := arpc.NewClientPool(func() (net.Conn, error) {
		return net.DialTimeout("tcp", listenAddr, time.Second*3)
	}, size)
	if err != nil {
		return nil, err
	}
	return &Client{pool: pool}, nil
}

// C 返回客户端实例
func (c *Client) C() *arpc.Client {
	return c.pool.Next()
}

package arpcclient

import (
	"sync"

	"github.com/lesismal/arpc"
)

var _clientOnce sync.Once
var _clientIns *Client

type Client struct {
	mu     sync.RWMutex
	pools  map[string]*arpc.ClientPool
	addrs  []string
	length int

	rk int

	cp ClientPool
}

// New 创建连接池
func New(listenAddr string, size int) (*Client, error) {
	_clientOnce.Do(func() {
		_clientIns = newClient()
	})

	if err := _clientIns.AddClientPool(listenAddr, size); err != nil {
		return nil, err
	}
	return _clientIns, nil
}

func newClient() *Client {
	pools := make(map[string]*arpc.ClientPool)
	return &Client{pools: pools, cp: NewDefaultClientPool()}
}

// AddClientPool 增加客户端池
func (c *Client) AddClientPool(listenAddr string, size int) error {
	pool, err := c.cp.Create(listenAddr, size)
	if err != nil {
		return err
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.pools[listenAddr] = pool
	c.addrs = append(c.addrs, listenAddr)
	c.length++
	return nil
}

// DeleteClientPool 删除客户端池
func (c *Client) DeleteClientPool(listenAddr string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.pools, listenAddr)
	keys := make([]string, len(c.pools))

	j := 0
	for k := range c.pools {
		keys[j] = k
		j++
	}
	c.addrs = keys
	c.length = j
}

// C 返回客户端实例
func (c *Client) C() *arpc.Client {
	c.mu.RLock()
	pool := c.pools[c.addrs[(c.rk+1)%c.length]]
	c.mu.RUnlock()
	return c.cp.Choose(pool)
}

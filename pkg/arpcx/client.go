package arpcx

import (
	"sync"

	"github.com/lesismal/arpc"
	"github.com/lesismal/arpc/codec"
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

	cc codec.Codec
}

// New 创建连接池
func NewClient() (*Client, error) {
	_clientOnce.Do(func() {
		_clientIns = newClient()
	})

	return _clientIns, nil
}

func newClient() *Client {
	pools := make(map[string]*arpc.ClientPool)
	return &Client{pools: pools, cp: NewDefaultClientPool()}
}

// SetCodec 设置客户端使用编码器
// - link: https://github.com/lesismal/arpc#custom-codec
func (c *Client) SetCodec(cc codec.Codec) {
	c.cc = cc
}

// Codec 获取设置的编码器;如果未设置，则 ok 为false
func (c *Client) Codec() (cc codec.Codec, ok bool) {
	if c.cc != nil {
		return c.cc, true
	}
	return nil, false
}

// AddNode 增加节点
// - listenAddr 以监听的节点地址作为节点映射
// - size 节点应该有的tcp连接数
func (c *Client) AddNode(listenAddr string, size int) error {
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

// DeleteNode 删除节点
func (c *Client) DeleteNode(listenAddr string) {
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

	arpcClient := c.cp.Choose(pool)

	// 初始化arpc.Client的步骤
	if cc, ok := c.Codec(); ok {
		arpcClient.Codec = cc
	}

	return arpcClient
}

package service

import (
	"fmt"
	"sync"

	"github.com/lesismal/arpc"
)

var ArcpClientMap sync.Map

// Arpc 抽象结构
type Arpc struct {
	Network string // 使用协议
	Host    string // ip
	Port    int    // 端口
	mu      sync.Mutex
}

type ArpcFactoryFunc func() (*arpc.Client, error)

// ArpcPool arpc连接池
type ArpcPool struct {
	mu      sync.Mutex
	res     chan *arpc.Client // 连接储存chan
	factory ArpcFactoryFunc   // 工厂方法
	closed  bool              // 关闭标识
}

// Acquire 获取一个连接
func (this *ArpcPool) Acquire() (*arpc.Client, error) {
	select {
	case c, ok := <-this.res:
		// 从池中获取
		if !ok {
			return nil, fmt.Errorf("pool closed!")
		}
		return c, nil
	default:
		// 新生成一个
		return this.factory()
	}
}

// Close 关闭连接池
func (this *ArpcPool) Close() {
	this.mu.Lock()
	defer this.mu.Unlock()

	if this.closed {
		return
	}

	this.closed = true
	close(this.res)
	for c := range this.res {
		c.Stop()
	}
}

// Release 释放资源，放回池中
func (this *ArpcPool) Release(c *arpc.Client) {
	if err := c.CheckState(); err != nil {
		// 检查Client的状态
		return
	}
	this.mu.Lock()
	defer this.mu.Unlock()

	// 只剩最后一个资源
	if this.closed {
		c.Stop()
		return
	}

	select {
	case this.res <- c:
		// 放回池中
	default:
		// 池子满了的情况
		c.Stop()
	}
}

// NewArpcPool 新建连接池
func NewArpcPool(fn ArpcFactoryFunc, size uint) (*ArpcPool, error) {
	if size <= 0 {
		return nil, fmt.Errorf("size is wrong!")
	}

	return &ArpcPool{
		factory: fn,
		res:     make(chan *arpc.Client, size),
	}, nil
}

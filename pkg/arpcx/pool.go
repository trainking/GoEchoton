package arpcx

import (
	"net"
	"time"

	"github.com/lesismal/arpc"
)

type (
	// ClientPool 抽象客户端池的操作
	ClientPool interface {
		//Create 创建一个用户池
		Create(listenAddr string, size int) (*arpc.ClientPool, error)

		// Choose 选择一个arpc.Client实例
		Choose(p *arpc.ClientPool) *arpc.Client
	}

	defaultClientPool struct{}
)

func NewDefaultClientPool() ClientPool {
	return &defaultClientPool{}
}

func (d *defaultClientPool) Create(listenAddr string, size int) (*arpc.ClientPool, error) {
	return arpc.NewClientPool(func() (net.Conn, error) {
		return net.DialTimeout("tcp", listenAddr, time.Second*5)
	}, size)
}

func (d *defaultClientPool) Choose(p *arpc.ClientPool) *arpc.Client {
	return p.Next()
}

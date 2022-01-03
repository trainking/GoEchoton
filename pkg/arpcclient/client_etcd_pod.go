package arpcclient

import (
	"GoEchoton/pkg/etcdx"

	"github.com/lesismal/arpc"
)

const DefaultPoolSize = 2

type ClientEtcdPod struct {
	*Client

	rempod *etcdx.RemotePod
	size   int
}

// NewClientPool 新建Etcd的客户端池
func NewClientPool(target string, etcdGateWay []string) (*ClientEtcdPod, error) {
	r, err := etcdx.NewRemotePod(target, etcdGateWay)
	if err != nil {
		return nil, err
	}
	c := &ClientEtcdPod{
		rempod: r,
		size:   DefaultPoolSize,
	}
	c.Client = newClient()

	// 增加变更回调
	c.rempod.SetOnAdd(func(v string) {
		if err := c.AddClientPool(v, c.size); err != nil {
			return
		}
	})
	c.rempod.SetOnDelete(func(v string) {
		c.DeleteClientPool(v)
	})

	return c, nil
}

// GetNode 获取一个节点
func (ce *ClientEtcdPod) GetNode() (*arpc.Client, error) {
	if ce.length == 0 {
		nodes := ce.rempod.GetNodes()
		if len(nodes) > 0 {
			for _, v := range nodes {
				if err := ce.AddClientPool(v, ce.size); err != nil {
					return nil, err
				}
			}
		}
	}
	return ce.C(), nil
}

// SetPoolSize 设置池的大小
func (ce *ClientEtcdPod) SetPoolSize(size int) {
	ce.size = size
}

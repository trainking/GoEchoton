package arpcx

import (
	"GoEchoton/pkg/etcdx"
	"context"

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

	c.Client.SetCodec(&MsgpackCodec{})

	// 增加变更回调
	c.rempod.SetOnAdd(func(v string) {
		if err := c.AddNode(v, c.size); err != nil {
			return
		}
	})
	c.rempod.SetOnDelete(func(v string) {
		c.DeleteNode(v)
	})

	// 初始化获取所有的节点
	if err := c.rempod.InitNodes(); err != nil {
		return nil, err
	}

	return c, nil
}

// GetNode 获取一个节点
func (ce *ClientEtcdPod) GetNode() *arpc.Client {
	if ce.length == 0 {
		nodes := ce.rempod.GetNodes()
		if len(nodes) > 0 {
			for _, v := range nodes {
				if err := ce.AddNode(v, ce.size); err != nil {
					panic(err)
				}
			}
		}
	}
	return ce.C()
}

func (ce *ClientEtcdPod) CallWith(ctx context.Context, method string, req interface{}, rsp interface{}) error {
	client := ce.GetNode()

	// 传递Request ID
	resuestID := ctx.Value("REQUEST_ID")
	if id, ok := resuestID.(string); ok {
		client.Set("REQUEST_ID", id)
	}

	return client.CallWith(ctx, method, req, rsp)
}

// SetPoolSize 设置池的大小
func (ce *ClientEtcdPod) SetPoolSize(size int) {
	ce.size = size
}

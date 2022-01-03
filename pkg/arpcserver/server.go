package arpcserver

import (
	"context"
	"fmt"
	"time"

	"github.com/lesismal/arpc"
	"github.com/lesismal/arpc/extension/middleware/router"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Server struct {
	server   *arpc.Server
	Handlers []Handler
}

func New(svcCtx ServerContext) *Server {
	return &Server{
		server:   arpc.NewServer(),
		Handlers: svcCtx.GetHandlers(),
	}
}

// RegisterToEtcd 注册服务到Etcd
func (s *Server) RegisterToEtcd(target string, value string, etcdGateway []string) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdGateway,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	kv := clientv3.NewKV(client)
	lease := clientv3.NewLease(client)
	var leaseID clientv3.LeaseID = 0

	go func() {
		defer client.Close()
		for {
			if leaseID == 0 {
				leaseResp, err := lease.Grant(context.TODO(), 10)
				if err != nil {
					panic(err)
				}
				key := target + fmt.Sprintf("%d", leaseResp.ID)
				if _, err := kv.Put(context.TODO(), key, value, clientv3.WithLease(leaseResp.ID)); err != nil {
					panic(err)
				}
				leaseID = leaseResp.ID
			} else {
				if _, err := lease.KeepAliveOnce(context.TODO(), leaseID); err == rpctypes.ErrLeaseNotFound {
					leaseID = 0
					continue
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()
}

// Start 启动
func (s *Server) Start(listenAddr string) {
	// 全局中间件
	s.server.Handler.Use(router.Recover())
	s.server.Handler.Use(router.Logger())

	// 加入调用
	for _, h := range s.Handlers {
		s.server.Handler.Handle(h.Path, Handle(h.Handle))
	}

	if err := s.server.Run(listenAddr); err != nil {
		panic(err)
	}
}

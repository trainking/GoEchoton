package arpcx

import (
	"GoEchoton/pkg/etcdx"
	"time"

	"github.com/lesismal/arpc"
	"github.com/lesismal/arpc/extension/middleware/router"
	"github.com/lesismal/arpc/log"
)

type Server struct {
	server   *arpc.Server
	Handlers []Handler
}

func NewServer(svcCtx ServerContext) *Server {
	return &Server{
		server:   arpc.NewServer(),
		Handlers: svcCtx.GetHandlers(),
	}
}

// RegisterToEtcd 注册服务到Etcd
func (s *Server) RegisterToEtcd(target string, value string, etcdGateway []string) {
	// 延迟注册，等待上一个租约过期，被客户端完全发现
	time.Sleep(3 * time.Second)

	// 注册到Etcd
	if err := etcdx.LeaseAndHeartbeat(target, value, etcdGateway, 3, 1); err != nil {
		log.Error(`Register Error: %s, Etcd: %v`, err.Error(), etcdGateway)
		panic(err)
	}
	log.Info(`Register Etcd: %v %s-%s`, etcdGateway, target, value)
}

// Start 启动
func (s *Server) Start(listenAddr string) {
	// 全局中间件
	s.server.Handler.Use(router.Recover())
	s.server.Handler.Use(router.Logger())
	s.server.Codec = &MsgpackCodec{}

	// 加入调用
	for _, h := range s.Handlers {
		s.server.Handler.Handle(h.Path, Handle(h.Handle))
	}

	if err := s.server.Run(listenAddr); err != nil {
		panic(err)
	}
}

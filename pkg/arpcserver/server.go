package arpcserver

import (
	"GoEchoton/pkg/etcdx"

	"github.com/lesismal/arpc"
	"github.com/lesismal/arpc/extension/middleware/router"
	"github.com/lesismal/arpc/log"
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
	if err := etcdx.LeaseAndHeartbeat(target, value, etcdGateway, 10, 1); err != nil {
		log.Error(`Rgister Error: %s, Etcd: %v`, err.Error(), etcdGateway)
		panic(err)
	}
	log.Info(`Register Etcd: %v %s-%s`, etcdGateway, target, value)
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

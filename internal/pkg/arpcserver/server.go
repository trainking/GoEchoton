package arpcserver

import (
	"github.com/lesismal/arpc"
	"github.com/lesismal/arpc/extension/middleware/router"
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

// Start 启动
func (s *Server) Strart(listenAddr string) {
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

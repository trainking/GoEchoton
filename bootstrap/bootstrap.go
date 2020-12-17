package bootstrap

import (
	. "GoEchoton/config"
	"GoEchoton/middleware"
	"GoEchoton/router"
	"fmt"

	"github.com/labstack/echo/v4"
)

type Server struct {
	e       *echo.Echo
	mFuns   []echo.MiddlewareFunc
	routers []router.Router
}

// 启动服务
func (s *Server) Start() {
	for _, f := range s.mFuns {
		s.e.Use(f)
	}
	for _, r := range s.routers {
		s.e.Add(r.Method, r.Path, r.Handler, r.Middlwares...)
	}
	s.e.Logger.Fatal(s.e.Start(fmt.Sprintf(":%d", Config.Server.Port)))
}

// 创建一个Server
func NewServer() *Server {
	var _e *echo.Echo = echo.New()
	var server *Server = &Server{
		e:       _e,
		mFuns:   middleware.Mfuns,
		routers: router.Routers,
	}
	return server
}

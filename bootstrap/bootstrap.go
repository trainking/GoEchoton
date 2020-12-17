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
	// 全局中间件
	for _, f := range s.mFuns {
		s.e.Use(f)
	}
	// 无组路由
	for _, r := range s.routers {
		s.e.Add(r.Method, r.Path, r.Handler, r.Middlwares...)
	}
	// 路由组
	for _, g := range router.Groups {
		_g := s.e.Group(g.Path)
		for _,_m := range g.Middlwares {
			_g.Use(_m)
		}
		for _,_r := range g.Routers {
			_g.Add(_r.Method, _r.Path, _r.Handler, _r.Middlwares...)
		}
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

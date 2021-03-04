package bootstrap

import (
	"GoEchoton/router"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server 服务结构体定义
type Server struct {
	e       *echo.Echo
	routers []router.Router
}

// Start 启动服务
func (s *Server) Start(port int) {
	// 全局中间件
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())

	// 字段验证器
	s.e.Validator = NewStructValidator()
	// Binder
	s.e.Binder = NewJsoniterBinder()
	// 无组路由
	for _, r := range s.routers {
		r.Add(s.e)
	}
	// 路由组
	for _, g := range router.Groups {
		g.Add(s.e)
	}
	s.e.Logger.Fatal(s.e.Start(fmt.Sprintf(":%d", port)))
}

// NewServer 创建一个Server
func NewServer() *Server {
	var _e *echo.Echo = echo.New()
	var server *Server = &Server{
		e:       _e,
		routers: router.Routers,
	}
	return server
}

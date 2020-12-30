package bootstrap

import (
	. "GoEchoton/config"
	"GoEchoton/router"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// StructValidator 结构体验证器
type StructValidator struct {
	validator *validator.Validate
}

// Validate 实现验证方法
func (s *StructValidator) Validate(i interface{}) error {
	return s.validator.Struct(i)
}

// NewStructValidator 创建验证器
func NewStructValidator() echo.Validator {
	return &StructValidator{
		validator: validator.New(),
	}
}

// Server 服务结构体定义
type Server struct {
	e       *echo.Echo
	routers []router.Router
}

// Start 启动服务
func (s *Server) Start() {
	// 全局中间件
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())

	// 字段验证器
	s.e.Validator = NewStructValidator()
	// 无组路由
	for _, r := range s.routers {
		r.Add(s.e)
	}
	// 路由组
	for _, g := range router.Groups {
		g.Add(s.e)
	}
	s.e.Logger.Fatal(s.e.Start(fmt.Sprintf(":%d", Config.Server.Port)))
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

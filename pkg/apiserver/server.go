package apiserver

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server api Server定义
type Server struct {
	e             *echo.Echo
	routers       []Router
	groups        []Group
	middlewares   []echo.MiddlewareFunc
	validator     *StructValidator
	ValidatorList map[string]validator.Func
}

// AddRouters 增加顶级路由
func (s *Server) AddRouters(l []Router) {
	if len(s.routers) == 0 {
		s.routers = l
	}
	s.routers = append(s.routers, l...)
}

// AddGroups 增加组路由
func (s *Server) AddGroups(l []Group) {
	if len(s.groups) == 0 {
		s.groups = l
	}
	s.groups = append(s.groups, l...)
}

// Start 开始服务，listenAddr 如 `127.0.0.1:5001`, `:5001`
func (s *Server) Start(listenAddr string) {
	// 加入验证器
	if len(s.ValidatorList) > 0 {
		for tag, f := range s.ValidatorList {
			if err := s.validator.AddValidator(tag, f); err != nil {
				panic(err)
			}
		}
	}
	s.e.Validator = s.validator.transEchoValidator()

	// 加入路由
	for _, g := range s.groups {
		g.Add(s.e)
	}
	for _, r := range s.routers {
		r.Add(s.e)
	}

	// 全局中间件
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())
	s.e.Use(middleware.CORS())
	s.e.Use(middleware.RequestID())
	// 额外增加的全局中间件
	for _, m := range s.middlewares {
		s.e.Use(m)
	}

	s.e.Logger.Fatal(s.e.Start(listenAddr))
}

// New 新建服务
func New(svcCtx ServerContext) *Server {
	return &Server{
		e:             echo.New(),
		routers:       svcCtx.GetRouters(),
		groups:        svcCtx.GetGroups(),
		middlewares:   svcCtx.GetMiddlewares(),
		ValidatorList: svcCtx.GetValidators(),
		validator:     NewStructValidator(),
	}
}

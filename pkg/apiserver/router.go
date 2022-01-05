package apiserver

import "github.com/labstack/echo/v4"

type (
	// Router 包装路由定义
	Router struct {
		Method     string
		Path       string
		Name       string
		Handler    HandlerFunc
		Middlwares []echo.MiddlewareFunc
	}

	// Group 路由组定义
	Group struct {
		Path       string
		Middlwares []echo.MiddlewareFunc
		Routers    []Router
	}
)

// Add 增加路由
func (r Router) Add(e *echo.Echo) {
	_r := e.Add(r.Method, r.Path, Handle(r.Handler), r.Middlwares...)
	if r.Name != "" {
		_r.Name = r.Name
	}
}

// Add 增加路由组
func (g Group) Add(e *echo.Echo) {
	_g := e.Group(g.Path)
	for _, _m := range g.Middlwares {
		_g.Use(_m)
	}
	for _, _r := range g.Routers {
		_rr := _g.Add(_r.Method, _r.Path, Handle(_r.Handler), _r.Middlwares...)
		if _r.Name != "" {
			_rr.Name = _r.Name
		}
	}
}

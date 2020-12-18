package router

import (
	"GoEchoton/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Router struct {
	Method     string
	Path       string
	Handler    echo.HandlerFunc
	Middlwares []echo.MiddlewareFunc
}

type Group struct {
	Path       string
	Middlwares []echo.MiddlewareFunc
	Routers    []Router
}

// 路由定义结构
var Routers []Router = []Router{}

// 分组路由
var Groups []Group = []Group{}

// 初始化加载路由
func init() {
	Routers = append(Routers, []Router{
		{
			Method:     http.MethodGet,
			Path:       "/",
			Handler:    handler.Index,
			Middlwares: nil,
		},
		{
			Method:     http.MethodPost,
			Path:       "/login",
			Handler:    handler.Login,
			Middlwares: nil,
		},
	}...)
}

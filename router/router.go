package router

import (
	"GoEchoton/handler"
	"net/http"
)

// 初始化加载路由
func init() {
	Routers = append(Routers, []Router{
		{
			Method:  http.MethodGet,
			Path:    "/whoami/",
			Handler: handler.User.Index,
		},
		{
			Method:  http.MethodPost,
			Path:    "/login",
			Handler: handler.User.Login,
		},
		{
			Method:  http.MethodGet,
			Path:    "/auth",
			Handler: handler.User.Auth,
		},
	}...)
}

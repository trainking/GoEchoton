package router

import (
	"GoEchoton/handler"
	"net/http"
)

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

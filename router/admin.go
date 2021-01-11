package router

import (
	. "GoEchoton/config"
	"GoEchoton/handler"
	"net/http"

	myware "GoEchoton/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	Groups = append(Groups, Group{
		Path: "/admin",
		Middlwares: []echo.MiddlewareFunc{
			middleware.JWT([]byte(Config.Jwt.Secret)),
			myware.Online(),
		},
		Routers: []Router{
			{
				Method:     http.MethodGet,
				Path:       "/index",
				Handler:    handler.User.Index,
				Middlwares: nil,
			},
		},
	})
}

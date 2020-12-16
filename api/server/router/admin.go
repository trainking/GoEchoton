package router

import (
	"GoEchoton/api/server/controller/admin"
	"GoEchoton/configs/api/conf"

	myware "GoEchoton/api/server/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func adminGroup(e *echo.Echo) {
	g := e.Group("/admin")
	g.Use(middleware.JWT([]byte(conf.Conf.Jwt.Secret)))
	g.Use(myware.Online())
	g.GET("/index", admin.Index)
}

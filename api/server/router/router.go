// @Titile /GoEchoton/api/router/router.go
// @Description api使用路由包
package router

import (
	"GoEchoton/api/server/controller/login"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @title Initized
// @description 注册路由函数
// @param e *echo.Echo  Echo对象
// @return error
func Initized(e *echo.Echo) error {

	e.GET("/", login.Hello)
	e.POST("/login", login.Login)
	e.GET("/routes", func(c echo.Context) error {
		return c.JSON(http.StatusOK, e.Routes())
	})

	// admin Group注册
	adminGroup(e)
	return nil
}

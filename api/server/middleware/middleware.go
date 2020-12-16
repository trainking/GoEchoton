// @Title /GoEchoton/api/middleware/middleware.go
// @Description api使用的中间件包
package middleware

import (
	"GoEchoton/api/types/hauthorized"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Online() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			t := c.Request().Header.Get(echo.HeaderAuthorization)
			r := hauthorized.Check(t)
			if !r {
				return &echo.HTTPError{
					Code:     http.StatusUnauthorized,
					Message:  "invalid or expired jwt",
					Internal: nil,
				}
			}
			return nil
		}
	}
}

// @title Initized
// @Description 注册中间件函数
func Initized() []echo.MiddlewareFunc {
	mFuns := []echo.MiddlewareFunc{}
	mFuns = append(mFuns, middleware.Logger())
	mFuns = append(mFuns, middleware.Recover())
	return mFuns
}

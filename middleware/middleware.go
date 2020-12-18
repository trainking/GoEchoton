package middleware

import (
	"GoEchoton/repository"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Online() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			t := c.Request().Header.Get(echo.HeaderAuthorization)
			var op repository.Hauthorized_OP = repository.NewHauthorizedOP()
			r := op.Check(t)
			if !r {
				return &echo.HTTPError{
					Code:     http.StatusUnauthorized,
					Message:  "invalid or expired jwt",
					Internal: nil,
				}
			}
			return next(c)
		}
	}
}

// 全局使用的中间件
var Mfuns []echo.MiddlewareFunc = []echo.MiddlewareFunc{}

func init() {
	Mfuns = append(Mfuns, middleware.Logger())
	Mfuns = append(Mfuns, middleware.Recover())
}

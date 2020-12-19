package middleware

import (
	"GoEchoton/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Online 在线状态中间件
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

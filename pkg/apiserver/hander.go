package apiserver

import "github.com/labstack/echo/v4"

// HandlerFunc 自定义HandlerFunc
type HandlerFunc func(ctx Context) error

func Handle(f HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := NewContext(c)
		return f(ctx)
	}
}

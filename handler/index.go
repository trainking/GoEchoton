package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// 首页Index
func Index(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"say": "hello, world!",
	})
}

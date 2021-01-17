package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Resp struct {
	Code int         `json:"code"` // 返回码
	Msg  string      `json:"msg"`  // 消息
	Data interface{} `json:"data"` // 数据
}

// ResponseSucessfully 成功返回
func ResponseSucessfully(c echo.Context, data interface{}) error {
	r := Resp{
		Code: 0,
		Msg:  "ok",
		Data: data,
	}
	return c.JSON(http.StatusOK, r)
}

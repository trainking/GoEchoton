package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	SuccessCode    = iota
	InnerErrorCode // 内部错误
)

// Resp 返回内容结构体
type Resp struct {
	Code int         `json:"code"` // 返回码
	Msg  string      `json:"msg"`  // 消息
	Data interface{} `json:"data"` // 数据
}

// Response 成功返回
func Response(c echo.Context, data interface{}) error {
	r := Resp{
		Code: SuccessCode,
		Msg:  "ok",
		Data: data,
	}
	return c.JSON(http.StatusOK, r)
}

// ErrorResponse 错误返回
func ErrorResponse(c echo.Context, msg string) error {
	r := Resp{
		Code: InnerErrorCode,
		Msg:  msg,
	}
	return c.JSON(http.StatusOK, r)
}

// ExtResponse 自定义返回
func ExtResponse(c echo.Context, r Resp) error {
	return c.JSON(http.StatusOK, r)
}

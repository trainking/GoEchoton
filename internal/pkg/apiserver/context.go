package apiserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var ResponseOkCode = 0   // 成功返回代码
var ResponseOkMsg = "ok" // 成功返回消息

type (
	Context interface {
		// BindAndValidate 绑定数据并验证
		BindAndValidate(i interface{}) error

		// GetCtx 获取Echo的context
		GetCtx() echo.Context

		// ErrResponse 错误返回
		ErrResponse(code int32, msg string) error

		// Response 正常返回
		Response(data interface{}) error
	}

	JsonResponse struct {
		Code int32       `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}

	defaultContext struct {
		ctx echo.Context
	}
)

// BindAndValidate 绑定数据并验证
func (c *defaultContext) BindAndValidate(i interface{}) error {
	if err := c.ctx.Bind(i); err != nil {
		return err
	}
	if err := c.ctx.Validate(i); err != nil {
		return err
	}
	return nil
}

// GetCtx 获取Echo的context
func (c *defaultContext) GetCtx() echo.Context {
	return c.ctx
}

// ErrResponse 错误返回 `{"code": 1, "msg": "error is ..."}`
func (c *defaultContext) ErrResponse(code int32, msg string) error {
	return c.ctx.JSON(http.StatusOK, JsonResponse{Code: code, Msg: msg})
}

// Response 正常返回 `{"code": 0, "msg": "ok", "data": ...}`
func (c *defaultContext) Response(data interface{}) error {
	return c.ctx.JSON(http.StatusOK, JsonResponse{Code: int32(ResponseOkCode), Msg: ResponseOkMsg})
}

// NewContext 创建默认context
func NewContext(ctx echo.Context) Context {
	return &defaultContext{
		ctx: ctx,
	}
}

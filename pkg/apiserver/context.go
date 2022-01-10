package apiserver

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const DefaultRequestTimout = 30 // 每个请求超时时间，最长30秒

var ResponseOkCode = 0   // 成功返回代码
var ResponseOkMsg = "ok" // 成功返回消息

var RequestID RequestIDT = "REQUEST_ID"

type (
	RequestIDT string

	Context interface {
		// BindAndValidate 绑定数据并验证
		BindAndValidate(i interface{}) error

		// GetCtx 获取Echo的context
		GetEchoContext() echo.Context

		// ErrResponse 错误返回
		ErrResponse(code int32, msg string) error

		// Response 正常返回
		Response(data interface{}) error

		// GetContext 获取一个context.Context
		GetContext() context.Context

		// Destoy 退出时操作
		Destoy()
	}

	JsonResponse struct {
		Code int32       `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}

	defaultContext struct {
		echoCtx echo.Context
		ctx     context.Context
		cancel  context.CancelFunc
	}
)

// BindAndValidate 绑定数据并验证
func (c *defaultContext) BindAndValidate(i interface{}) error {
	if err := c.echoCtx.Bind(i); err != nil {
		return err
	}
	if err := c.echoCtx.Validate(i); err != nil {
		return err
	}
	return nil
}

// GetCtx 获取Echo的context
func (c *defaultContext) GetEchoContext() echo.Context {
	return c.echoCtx
}

// ErrResponse 错误返回 `{"code": 1, "msg": "error is ..."}`
func (c *defaultContext) ErrResponse(code int32, msg string) error {
	return c.echoCtx.JSON(http.StatusOK, JsonResponse{Code: code, Msg: msg})
}

// Response 正常返回 `{"code": 0, "msg": "ok", "data": ...}`
func (c *defaultContext) Response(data interface{}) error {
	return c.echoCtx.JSON(http.StatusOK, JsonResponse{Code: int32(ResponseOkCode), Msg: ResponseOkMsg, Data: data})
}

//GetContext 获取一个context.Context
func (c *defaultContext) GetContext() context.Context {
	return c.ctx
}

// Destoy 退出时应做的操作
func (c *defaultContext) Destoy() {
	c.cancel()
}

// NewContext 创建默认context
func NewContext(echoCtx echo.Context) Context {
	ctx := context.WithValue(context.Background(), RequestID, echoCtx.Request().Header.Get(echo.HeaderXRequestID))
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, DefaultRequestTimout*time.Second)
	return &defaultContext{
		echoCtx: echoCtx,
		ctx:     ctx,
		cancel:  cancel,
	}
}

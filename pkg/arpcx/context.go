package arpcx

import (
	"context"

	"github.com/lesismal/arpc"
)

type (
	Context interface {

		// GetRequestID 获取RequestID
		GetRequestID() string

		// Bind 将数据绑定
		Bind(i interface{}) error

		// Write 写入成功返回
		Write(i ...interface{}) error

		// GetContext 获取一个context.Context
		GetContext() context.Context
	}

	defaultContext struct {
		ctx *arpc.Context
	}
)

func NewContext(ctx *arpc.Context) Context {
	return &defaultContext{
		ctx: ctx,
	}
}

// GetRequestID 获取RequestID
func (c *defaultContext) GetRequestID() string {
	if id, ok := c.ctx.Get("REQUEST_ID"); ok {
		if _ids, ok := id.(string); ok {
			return _ids
		}
	}
	return ""
}

// Bind 绑定数据
func (c *defaultContext) Bind(i interface{}) error {
	return c.ctx.Bind(i)
}

// Write 写出返回结果
func (c *defaultContext) Write(i ...interface{}) error {
	if len(i) == 0 {
		return c.ctx.Write(struct{}{})
	}
	if len(i) > 1 {
		panic("write more than one")
	}
	return c.ctx.Write(i[0])
}

// GetContext 获取一个context.Context
func (c *defaultContext) GetContext() context.Context {
	return c.ctx
}

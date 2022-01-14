package arpcx

import (
	"encoding/json"

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
	}

	Request struct {
		ID   string      `json:"id"`
		Data interface{} `json:"data"`
	}

	defaultContext struct {
		ctx *arpc.Context
		req *Request
	}
)

func NewContext(ctx *arpc.Context) Context {
	var req Request
	if err := ctx.Bind(&req); err != nil {
		panic(err)
	}
	return &defaultContext{
		ctx: ctx,
		req: &req,
	}
}

// GetRequestID 获取RequestID
func (c *defaultContext) GetRequestID() string {
	return c.req.ID
}

// Bind 绑定数据
func (c *defaultContext) Bind(i interface{}) error {
	b, err := json.Marshal(c.req.Data)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, i)
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

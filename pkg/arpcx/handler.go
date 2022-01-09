package arpcx

import (
	"github.com/lesismal/arpc"
	"github.com/lesismal/arpc/log"
)

type Result struct {
	Err  error
	Data interface{}
}

// HandlerFuc Handler
type HandlerFunc func(ctx *arpc.Context) *Result

// Handle 转换Handler
func Handle(f HandlerFunc) arpc.HandlerFunc {
	return func(ctx *arpc.Context) {
		_r := f(ctx)
		if _r.Err != nil {
			// 记录返回错误日志
			method := ctx.Message.Method()
			addr := ctx.Client.Conn.RemoteAddr()
			log.Error(`%s %v Error: %s`, method, addr, _r.Err.Error())

			if err := ctx.Error(_r.Err); err != nil {
				panic(err)
			}
		} else {
			if _r.Data == nil {
				_r.Data = struct{}{}
			}
			if err := ctx.Write(_r.Data); err != nil {
				panic(err)
			}
		}
	}
}

// Handler 定义Handler
type Handler struct {
	Path   string
	Handle HandlerFunc
}

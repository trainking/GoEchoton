package arpcx

import (
	"github.com/lesismal/arpc"
	"github.com/lesismal/arpc/log"
)

// HandlerFuc Handler
// type HandlerFunc func(ctx *arpc.Context) *Result
type HandlerFunc func(ctx Context) error

// Handle 转换Handler
func Handle(f HandlerFunc) arpc.HandlerFunc {
	return func(c *arpc.Context) {
		ctx := NewContext(c)
		err := f(ctx)
		if err != nil {
			method := c.Message.Method()
			addr := c.Client.Conn.RemoteAddr()
			log.Error(`%s %v Error: %s`, method, addr, err.Error())
			if err := c.Error(err); err != nil {
				panic(err)
			}
			return
		}
	}
}

// Handler 定义Handler
type Handler struct {
	Path   string
	Handle HandlerFunc
}

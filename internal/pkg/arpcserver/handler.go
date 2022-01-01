package arpcserver

import "github.com/lesismal/arpc"

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
			if err := ctx.Error(_r.Err); err != nil {
				panic(err)
			}
		} else {
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

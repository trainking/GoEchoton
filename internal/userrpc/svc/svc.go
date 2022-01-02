package svc

import (
	"GoEchoton/internal/pkg/arpcserver"
	"GoEchoton/internal/userrpc/types"

	"github.com/lesismal/arpc"
)

type SvcContext struct {
}

func New() arpcserver.ServerContext {
	return &SvcContext{}
}

// GetHandlers 定义的Handler
func (s *SvcContext) GetHandlers() []arpcserver.Handler {
	return []arpcserver.Handler{
		{
			Path: types.UserCheckPasswdPath,
			Handle: func(ctx *arpc.Context) *arpcserver.Result {
				var p types.CheckPasswdApply
				if err := ctx.Bind(&p); err != nil {
					return &arpcserver.Result{Err: err}
				}
				return &arpcserver.Result{}
			},
		},
	}
}

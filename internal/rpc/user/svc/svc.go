package svc

import (
	"GoEchoton/internal/rpc/user/skeleton/handler"
	"GoEchoton/internal/rpc/user/userrpc"
	"GoEchoton/pkg/arpcx"
)

type SvcContext struct {
}

func New() arpcx.ServerContext {
	return &SvcContext{}
}

// GetHandlers 定义的Handler
func (s *SvcContext) GetHandlers() []arpcx.Handler {
	return []arpcx.Handler{
		{
			Path:   userrpc.UserCheckPasswdPath,
			Handle: handler.UserCheckPasswd,
		},
	}
}

package svc

import (
	"GoEchoton/internal/rpc/user/skeleton/handler"
	"GoEchoton/internal/rpc/user/userrpc"
	"GoEchoton/pkg/arpcserver"
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
			Path:   userrpc.UserCheckPasswdPath,
			Handle: handler.UserCheckPasswd,
		},
	}
}

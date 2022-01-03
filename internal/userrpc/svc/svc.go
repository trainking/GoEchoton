package svc

import (
	"GoEchoton/internal/userrpc/handler"
	"GoEchoton/internal/userrpc/types"
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
			Path:   types.UserCheckPasswdPath,
			Handle: handler.UserCheckPasswd,
		},
	}
}

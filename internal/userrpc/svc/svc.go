package svc

import (
	"GoEchoton/internal/pkg/arpcserver"
	"GoEchoton/internal/userrpc/handler"
	"GoEchoton/internal/userrpc/types"
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

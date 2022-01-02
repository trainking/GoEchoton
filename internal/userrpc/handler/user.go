package handler

import (
	"GoEchoton/internal/pkg/arpcserver"
	"GoEchoton/internal/userrpc/service"
	"GoEchoton/internal/userrpc/types"
	"context"

	"github.com/lesismal/arpc"
)

// UserCheckPasswd 检查密码
func UserCheckPasswd(ctx *arpc.Context) *arpcserver.Result {
	var p types.CheckPasswd
	if err := ctx.Bind(&p); err != nil {
		return &arpcserver.Result{Err: err}
	}
	if err := service.NewUserService().CheckPasswd(context.Background(), &p); err != nil {
		return &arpcserver.Result{Err: err}
	}
	return &arpcserver.Result{}
}

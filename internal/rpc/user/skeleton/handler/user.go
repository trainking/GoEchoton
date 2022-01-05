package handler

import (
	"GoEchoton/internal/rpc/user/skeleton/service"
	"GoEchoton/internal/rpc/user/userrpc"
	"GoEchoton/pkg/arpcserver"
	"context"

	"github.com/lesismal/arpc"
)

// UserCheckPasswd 检查密码
func UserCheckPasswd(ctx *arpc.Context) *arpcserver.Result {
	var p userrpc.CheckPasswd
	if err := ctx.Bind(&p); err != nil {
		return &arpcserver.Result{Err: err}
	}
	if err := service.NewUserService().CheckPasswd(context.Background(), &p); err != nil {
		return &arpcserver.Result{Err: err}
	}
	return &arpcserver.Result{}
}

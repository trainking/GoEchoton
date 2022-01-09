package handler

import (
	"GoEchoton/internal/rpc/user/skeleton/service"
	"GoEchoton/internal/rpc/user/userrpc"
	"GoEchoton/pkg/arpcx"
	"context"

	"github.com/lesismal/arpc"
)

// UserCheckPasswd 检查密码
func UserCheckPasswd(ctx *arpc.Context) *arpcx.Result {
	var p userrpc.CheckPasswd
	if err := ctx.Bind(&p); err != nil {
		return &arpcx.Result{Err: err}
	}
	if err := service.NewUserService().CheckPasswd(context.Background(), &p); err != nil {
		return &arpcx.Result{Err: err}
	}
	return &arpcx.Result{}
}

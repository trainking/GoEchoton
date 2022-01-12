package handler

import (
	"GoEchoton/internal/rpc/user/skeleton/service"
	"GoEchoton/internal/rpc/user/userrpc"
	"GoEchoton/pkg/arpcx"
	"context"
)

// UserCheckPasswd 检查密码
func UserCheckPasswd(ctx arpcx.Context) error {
	var p userrpc.CheckPasswd
	if err := ctx.Bind(&p); err != nil {
		return err
	}

	if err := service.NewUserService().CheckPasswd(context.Background(), &p); err != nil {
		return err
	}

	return ctx.Write()
}

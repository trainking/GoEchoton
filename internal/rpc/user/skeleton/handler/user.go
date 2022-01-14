package handler

import (
	"GoEchoton/internal/rpc/user/skeleton/service"
	"GoEchoton/internal/rpc/user/userrpc"
	"GoEchoton/pkg/arpcx"
)

// UserCheckPasswd 检查密码
func UserCheckPasswd(ctx arpcx.Context) error {
	var p userrpc.CheckPasswd
	if err := ctx.Bind(&p); err != nil {
		return err
	}

	if err := service.NewUserService().CheckPasswd(ctx.GetContext(), &p); err != nil {
		return err
	}

	return ctx.Write()
}

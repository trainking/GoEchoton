package login

import (
	"GoEchoton/internal/rpc/user/userrpc"
	"GoEchoton/internal/server/auth/apply"
	"GoEchoton/internal/server/auth/reply"
	"GoEchoton/pkg/apiserver"
	"context"
)

type (
	LoginApi interface {
		// LoginOne 登录第一步
		LoginOne(ctx apiserver.Context) error
	}

	loginApi struct {
		userRpc userrpc.UserRpc
	}
)

func New(userRpc userrpc.UserRpc) LoginApi {
	return &loginApi{
		userRpc: userRpc,
	}
}

//LoginOne 登录第一步
func (l *loginApi) LoginOne(ctx apiserver.Context) error {
	var p apply.LoginOneApply

	if err := ctx.BindAndValidate(&p); err != nil {
		return ctx.ErrResponse(1, err.Error())
	}

	// 检查密码
	if err := l.userRpc.CheckPasswd(context.Background(), &userrpc.CheckPasswd{
		Account:  p.Account,
		Password: p.Password,
	}); err != nil {
		return ctx.ErrResponse(1, err.Error())
	}

	return ctx.Response(reply.LoginOneReply{})
}

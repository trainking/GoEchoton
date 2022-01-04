package login

import (
	"GoEchoton/internal/authserver/apply"
	"GoEchoton/internal/authserver/reply"
	"GoEchoton/internal/userrpc/types"
	"GoEchoton/pkg/apiserver"
	"context"

	"github.com/labstack/echo/v4"
)

type (
	LoginApi interface {
		// LoginOne 登录第一步
		LoginOne(c echo.Context) error
	}

	loginApi struct {
		userRpc types.UserRpc
	}
)

func New(userRpc types.UserRpc) LoginApi {
	return &loginApi{
		userRpc: userRpc,
	}
}

//LoginOne 登录第一步
func (l *loginApi) LoginOne(c echo.Context) error {
	ctx := apiserver.NewContext(c)
	var p apply.LoginOneApply

	if err := ctx.BindAndValidate(&p); err != nil {
		return ctx.ErrResponse(1, err.Error())
	}

	// 检查密码
	if err := l.userRpc.CheckPasswd(context.Background(), &types.CheckPasswd{
		Account:  p.Account,
		Password: p.Password,
	}); err != nil {
		return ctx.ErrResponse(1, err.Error())
	}

	return ctx.Response(reply.LoginOneReply{})
}

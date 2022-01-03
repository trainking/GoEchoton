package login

import (
	"GoEchoton/internal/authserver/apply"
	"GoEchoton/internal/authserver/reply"
	"GoEchoton/internal/userrpc/types"
	"GoEchoton/internal/userrpc/userclient"
	"GoEchoton/pkg/apiserver"
	"context"

	"github.com/labstack/echo/v4"
)

type Login struct {
}

func New() *Login {
	return &Login{}
}

//LoginOne 登录第一步
func (l *Login) LoginOne(c echo.Context) error {
	ctx := apiserver.NewContext(c)
	var p apply.LoginOneApply

	if err := ctx.BindAndValidate(&p); err != nil {
		return ctx.ErrResponse(1, err.Error())
	}

	// 检查密码
	if err := userclient.NewUserRpc("127.0.0.1:8080").CheckPasswd(context.Background(), &types.CheckPasswd{
		Account:  p.Account,
		Password: p.Password,
	}); err != nil {
		return ctx.ErrResponse(1, err.Error())
	}

	return ctx.Response(reply.LoginOneReply{})
}

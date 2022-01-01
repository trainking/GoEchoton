package login

import (
	"GoEchoton/internal/authserver/types"
	"GoEchoton/internal/pkg/apiserver"

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
	var p types.LoginOneApply

	if err := ctx.BindAndValidate(&p); err != nil {
		return ctx.ErrResponse(1, err.Error())
	}

	return ctx.Response(types.LoginOneReply{})
}

package svc

import (
	"GoEchoton/internal/authserver/api/login"
	"GoEchoton/pkg/apiserver"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"honnef.co/go/tools/config"
)

type SvcContext struct {
	conf *config.Config
}

func New(conf *config.Config) apiserver.ServerContext {
	return &SvcContext{
		conf: conf,
	}
}

//GetRouters 获取顶级路由
func (s *SvcContext) GetRouters() []apiserver.Router {
	loginApi := login.New()
	return []apiserver.Router{
		{
			Method:  http.MethodPost,
			Path:    "/v1/login/one",
			Name:    "登录第一步",
			Handler: loginApi.LoginOne,
		},
	}
}

// GetGroups 获取组路由
func (s *SvcContext) GetGroups() []apiserver.Group {
	return []apiserver.Group{}
}

func (s *SvcContext) GetMiddlewares() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{}
}

// GetValidators 获取自定义验证器
func (s *SvcContext) GetValidators() map[string]validator.Func {
	return map[string]validator.Func{}
}

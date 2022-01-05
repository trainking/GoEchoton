package svc

import (
	"GoEchoton/internal/rpc/user/userstub"
	"GoEchoton/internal/server/auth/api/login"
	"GoEchoton/pkg/apiserver"
	"net/http"

	"GoEchoton/internal/server/auth/config"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

const AuthServerConfigEtcdPath = "/authserver/config"

type SvcContext struct {
	conf        *config.Config
	etcdGateway []string

	// api 对象
	loginApi login.LoginApi
}

func New(conf *config.Config, etcdGateway []string) apiserver.ServerContext {
	return &SvcContext{
		conf:        conf,
		etcdGateway: etcdGateway,

		loginApi: login.New(userstub.NewUserRpc(etcdGateway)),
	}
}

//GetRouters 获取顶级路由
func (s *SvcContext) GetRouters() []apiserver.Router {
	return []apiserver.Router{
		{
			Method:  http.MethodPost,
			Path:    "/v1/login/one",
			Name:    "登录第一步",
			Handler: s.loginApi.LoginOne,
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

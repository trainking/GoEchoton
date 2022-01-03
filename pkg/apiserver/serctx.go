package apiserver

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ServerContext interface {
	// GetRouters 获取顶级路由
	GetRouters() []Router

	// GetGroups 获取组路由
	GetGroups() []Group

	// GetMiddlewares 获取全局路由
	GetMiddlewares() []echo.MiddlewareFunc

	// GetValidators 获取全局验证器
	GetValidators() map[string]validator.Func
}

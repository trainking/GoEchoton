package apiserver

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// StructValidator 结构体验证器
type StructValidator struct {
	validator *validator.Validate
}

// Validate 实现echo.Validator
func (s *StructValidator) Validate(i interface{}) error {
	return s.validator.Struct(i)
}

// AddValidator 增加自定义验证器
func (s *StructValidator) AddValidator(tag string, v validator.Func) error {
	return s.validator.RegisterValidation(tag, v)
}

// transEchoValidator 转换为echo接口
func (s *StructValidator) transEchoValidator() echo.Validator {
	return s
}

// NewStructValidator 构建新结构体
func NewStructValidator() *StructValidator {
	return &StructValidator{validator: validator.New()}
}

package bootstrap

import (
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// StructValidator 结构体验证器
type StructValidator struct {
	validator *validator.Validate
}

// Validate 实现验证方法
func (s *StructValidator) Validate(i interface{}) error {
	return s.validator.Struct(i)
}

// idValid id验证器
func idValid(fl validator.FieldLevel) bool {
	switch fl.Field().Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint64:
		return true
	default:
		return regexp.MustCompile(`^[1-9]\d*$`).MatchString(fl.Field().String())
	}
}

// NewStructValidator 创建验证器
func NewStructValidator() echo.Validator {
	validator := validator.New()

	// 添加自定义验证器
	validator.RegisterValidation("idvalid", idValid)
	return &StructValidator{
		validator: validator,
	}
}

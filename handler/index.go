package handler

import (
	. "GoEchoton/config"
	"GoEchoton/repository"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// Index 首页Index
func Index(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"say": "hello, world!",
	})
}

// Login 登陆
func Login(c echo.Context) error {
	var param struct {
		Username string
		Password string
	}
	c.Bind(&param)
	userop := repository.NewUserOP()
	_r, err := userop.Valid(param.Username, param.Password)
	if err != nil {
		return err
	}
	if !_r {
		return echo.ErrUnauthorized
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Jon Snow"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(Config.Jwt.Secret))
	if err != nil {
		return err
	}
	op, err := repository.NewHauthorizedOP()
	if err != nil {
		return err
	}
	err = op.Save(param.Username, "Bearer "+t)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

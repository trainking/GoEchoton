package handler

import (
	. "GoEchoton/config"
	"GoEchoton/model/param"
	"GoEchoton/repository"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type user struct{}

var User = user{}

// Index 首页Index
func (_ user) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"say": "hello, world!",
	})
}

// Login 登陆
func (_ user) Login(c echo.Context) error {
	var params param.LoginUser
	c.Bind(&params)
	userop := repository.NewUserOP()
	_r, err := userop.Valid(params.Username, params.Password)
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
	err = op.Save(params.Username, "Bearer "+t)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

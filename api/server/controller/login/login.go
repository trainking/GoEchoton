package login

import (
	"GoEchoton/api/types/hauthorized"
	"GoEchoton/configs/api/conf"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// 判断用户名和密码是否一致
	if username != "jon" || password != "hahha" {
		return echo.ErrUnauthorized
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Jon Snow"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(conf.Conf.Jwt.Secret))
	if err != nil {
		return err
	}
	err = hauthorized.Save(username, "Bearer "+t)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

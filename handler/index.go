package handler

import (
	"GoEchoton/config"
	"GoEchoton/repository"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

// 首页Index
func Index(c echo.Context) error {
	fmt.Print("sss")
	return c.JSON(http.StatusOK, map[string]string{
		"say": "hello, world!",
	})
}

// 登录传参数结构体
type Login_json struct {
	Username string
	Password string
}

var ctx = context.Background()

// 验证账号和密码
func (l *Login_json) valid() (bool, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "121.4.65.13:6379",
		Password: "y123456", // no password set
		DB:       0,         // use default DB
	})
	var username string
	var password string
	var err error
	username, err = rdb.Get(ctx, "username").Result()
	if err != nil {
		return false, err
	}
	password, err = rdb.Get(ctx, "password").Result()
	if err != nil {
		return false, err
	}
	if l.Username != username || l.Password != password {
		return false, nil
	}
	return true, nil
}

// 登陆
func Login(c echo.Context) error {
	var param Login_json
	c.Bind(&param)
	_r, err := param.valid()
	if !_r {
		return echo.ErrUnauthorized
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Jon Snow"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(config.Config.Jwt.Secret))
	if err != nil {
		return err
	}
	op := repository.NewHauthorizedOP()
	err = op.Save(param.Username, "Bearer "+t)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

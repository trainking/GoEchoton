package handler

import (
	. "GoEchoton/config"
	"GoEchoton/repository"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

// Index 首页Index
func Index(c echo.Context) error {
	fmt.Print("sss")
	return c.JSON(http.StatusOK, map[string]string{
		"say": "hello, world!",
	})
}

// Login_json 登录传参数结构体
type Login_json struct {
	Username string
	Password string
}

// valid 验证账号和密码
func (l *Login_json) valid() (bool, error) {
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     Config.Redis.Host,
		Password: Config.Redis.Passwd,
		DB:       Config.Redis.DB,
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

// Login 登陆
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

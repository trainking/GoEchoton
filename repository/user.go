package repository

import (
	"GoEchoton/database"
)

type (
	// UserOP 用户操作接口
	UserOP interface {
		Repository

		// Valid 校验用户名和密码
		Valid(username, password string) (bool, error)
	}

	// userop op实现
	userop struct {
		client *database.Redis
	}
)

// Valid 校验用户名和密码
func (op *userop) Valid(username, password string) (bool, error) {
	uname, err := op.client.Get("username")
	if err != nil {
		return false, err
	}
	upasswd, err := op.client.Get("password")
	if err != nil {
		return false, err
	}
	if uname != username || upasswd != password {
		return false, nil
	}
	return true, nil
}

// Destory
func (op *userop) Destory() error {
	return nil
}

// NewUserOP 创建OP
func NewUserOP() UserOP {
	return &userop{
		client: database.NewRedis(),
	}
}

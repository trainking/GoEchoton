package service

import (
	"GoEchoton/internal/userrpc/types"
	"context"
	"sync"
)

var _userServiceOnce sync.Once
var _userServiceIns *UserService

type UserService struct {
}

func NewUserService() types.UserRpc {
	_userServiceOnce.Do(func() {
		_userServiceIns = &UserService{}
	})
	return _userServiceIns
}

// CheckPasswd 检查密码
func (s *UserService) CheckPasswd(ctx context.Context, p *types.CheckPasswd) error {
	// TODO 业务代码
	return nil
}

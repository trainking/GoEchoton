package types

import "context"

type (

	// UserRpc 定义接口
	UserRpc interface {
		//CheckPasswd 检查密码
		CheckPasswd(ctx context.Context, p *CheckPasswdApply) error
	}

	CheckPasswdApply struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}
)

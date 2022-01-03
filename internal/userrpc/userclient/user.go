package userclient

import (
	"GoEchoton/internal/userrpc/types"
	"GoEchoton/pkg/arpcclient"
	"context"
)

type userRpc struct {
	client *arpcclient.Client
}

func NewUserRpc(listenAddr string) types.UserRpc {
	client, err := arpcclient.New(listenAddr, 2)
	if err != nil {
		panic(err)
	}
	return &userRpc{
		client: client,
	}
}

// CheckPasswd 检查密码
func (u *userRpc) CheckPasswd(ctx context.Context, p *types.CheckPasswd) error {
	return u.client.C().CallWith(ctx, types.UserCheckPasswdPath, p, &struct{}{})
}

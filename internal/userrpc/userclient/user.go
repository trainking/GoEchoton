package userclient

import (
	"GoEchoton/internal/userrpc/types"
	"GoEchoton/pkg/arpcclient"
	"context"
)

type userRpc struct {
	client *arpcclient.ClientEtcdPod
}

func NewUserRpc(etcdGateway []string) types.UserRpc {
	client, err := arpcclient.NewClientPool(types.UserRpcTarget, etcdGateway)
	if err != nil {
		panic(err)
	}
	return &userRpc{
		client: client,
	}
}

// CheckPasswd 检查密码
func (u *userRpc) CheckPasswd(ctx context.Context, p *types.CheckPasswd) error {
	return u.client.GetNode().CallWith(ctx, types.UserCheckPasswdPath, p, &struct{}{})
}

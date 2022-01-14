package userstub

import (
	"GoEchoton/internal/rpc/user/userrpc"
	"GoEchoton/pkg/arpcx"
	"context"
)

type userRpc struct {
	client *arpcx.ClientEtcdPod
}

func NewUserRpc(etcdGateway []string) userrpc.UserRpc {
	client, err := arpcx.NewClientPool(userrpc.UserRpcTarget, etcdGateway)
	if err != nil {
		panic(err)
	}
	return &userRpc{
		client: client,
	}
}

// CheckPasswd 检查密码
func (u *userRpc) CheckPasswd(ctx context.Context, p *userrpc.CheckPasswd) error {
	return u.client.CallWith(ctx, userrpc.UserCheckPasswdPath, p, &struct{}{})
}

# arpcclient

## 概述

基于`arpc`抽象出来的rpc客户端

## 快速开始

```go
package userclient

import (
	"GoEchoton/internal/rpc/user/userrpc"
	"GoEchoton/pkg/arpcclient"
	"context"
)

type userRpc struct {
	client *arpcclient.ClientEtcdPod
}

func NewUserRpc(etcdGateway []string) userrpc.UserRpc {
	client, err := arpcclient.NewClientPool(userrpc.UserRpcTarget, etcdGateway)
	if err != nil {
		panic(err)
	}
	return &userRpc{
		client: client,
	}
}

// CheckPasswd 检查密码
func (u *userRpc) CheckPasswd(ctx context.Context, p *userrpc.CheckPasswd) error {
	return u.client.GetNode().CallWith(ctx, userrpc.UserCheckPasswdPath, p, &struct{}{})
}

```
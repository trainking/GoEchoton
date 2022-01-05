package main

import (
	"GoEchoton/internal/rpc/user/svc"
	"GoEchoton/internal/rpc/user/userrpc"
	"GoEchoton/pkg/arpcserver"
	"flag"
	"strings"
)

var listenAddr = flag.String("addr", ":8001", "listen address")
var etcdGateway = flag.String("etcd", "127.0.0.1:2379", "etcdGateway")

func main() {
	flag.Parse()
	svcCtx := svc.New()
	server := arpcserver.New(svcCtx)

	// 注册服务
	server.RegisterToEtcd(userrpc.UserRpcTarget, *listenAddr, strings.Split(*etcdGateway, ","))

	server.Start(*listenAddr)
}

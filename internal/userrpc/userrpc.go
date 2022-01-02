package main

import (
	"GoEchoton/internal/pkg/arpcserver"
	"GoEchoton/internal/userrpc/svc"
	"GoEchoton/internal/userrpc/types"
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
	server.RegisterToEtcd(types.UserRpcTarget, *listenAddr, strings.Split(*etcdGateway, ","))

	server.Start(*listenAddr)
}

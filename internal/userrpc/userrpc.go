package main

import (
	"GoEchoton/internal/pkg/arpcserver"
	"GoEchoton/internal/userrpc/svc"
	"flag"
)

var listenAddr = flag.String("addr", ":8001", "listen address")

func main() {
	flag.Parse()
	svcCtx := svc.New()
	server := arpcserver.New(svcCtx)

	// 注册服务
	server.RegisterToEtcd("/user.rpc/", *listenAddr, []string{"127.0.0.1:2379"})

	server.Start(*listenAddr)
}

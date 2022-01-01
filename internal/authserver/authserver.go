package main

import (
	"GoEchoton/internal/authserver/svc"
	"GoEchoton/internal/pkg/apiserver"
	"flag"

	"honnef.co/go/tools/config"
)

// 监听地址
var listenAddr = flag.String("addr", ":5050", "listen address")

func main() {
	flag.Parse()

	svcCtx := svc.New(&config.Config{})
	server := apiserver.New(svcCtx)
	server.Start(*listenAddr)
}

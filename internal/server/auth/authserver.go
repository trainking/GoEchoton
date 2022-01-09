package main

import (
	"GoEchoton/internal/server/auth/svc"
	"GoEchoton/pkg/apiserver"
	"GoEchoton/pkg/etcdx"
	"flag"
	"strings"

	"GoEchoton/internal/server/auth/config"
)

// 监听地址
var listenAddr = flag.String("addr", ":5050", "listen address")
var etcdGateway = flag.String("etcd", "127.0.0.1:8001", "etcd gateway")

func main() {
	flag.Parse()

	// 加载配置
	var conf config.Config
	var gateway = strings.Split(*etcdGateway, ",")
	clientX := etcdx.New(gateway)
	if err := apiserver.LoadConfigEtcdX(svc.AuthServerConfigEtcdPath, clientX, &conf); err != nil {
		panic(err)
	}

	svcCtx := svc.New(&conf, gateway)
	server := apiserver.New(svcCtx)
	server.Start(*listenAddr)
}

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

	var conf config.Config
	var gateway = strings.Split(*etcdGateway, ",")
	clientx := etcdx.New(gateway)
	if err := apiserver.LoadConfigEtcdX(svc.AuthServerConfigEtcdPath, clientx, &conf); err != nil {
		panic(err)
	}

	svcCtx := svc.New(&conf, gateway)
	server := apiserver.New(svcCtx)
	server.Start(*listenAddr)
}

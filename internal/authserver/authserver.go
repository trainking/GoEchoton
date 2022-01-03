package main

import (
	"GoEchoton/internal/authserver/svc"
	"GoEchoton/pkg/apiserver"
	"flag"
	"strings"

	"honnef.co/go/tools/config"
)

// 监听地址
var listenAddr = flag.String("addr", ":5050", "listen address")
var etcdGateway = flag.String("etcd", "127.0.0.1:8001", "etcd gateway")

func main() {
	flag.Parse()

	var conf config.Config
	// if err := apiserver.LoadConfigFile("./configs/authserver.yaml", &conf); err != nil {
	// 	panic(err)
	// }
	svcCtx := svc.New(&conf, strings.Split(*etcdGateway, ","))
	server := apiserver.New(svcCtx)
	server.Start(*listenAddr)
}

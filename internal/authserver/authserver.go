package main

import (
	"GoEchoton/internal/pkg/apiserver"
	"flag"
)

// 监听地址
var listenAddr = flag.String("addr", ":5050", "listen address")

func main() {
	flag.Parse()

	server := apiserver.New()
	server.Start(*listenAddr)
}

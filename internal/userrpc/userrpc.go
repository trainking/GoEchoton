package main

import (
	"GoEchoton/internal/pkg/arpcserver"
	"GoEchoton/internal/userrpc/svc"
	"flag"
)

var listenAddr = flag.String("addr", ":8080", "listen address")

func main() {
	flag.Parse()
	svcCtx := svc.New()
	server := arpcserver.New(svcCtx)

	server.Strart(*listenAddr)
}

package main

import (
	"GoEchoton/bootstrap"
	. "GoEchoton/config"
	"flag"
)

func main() {

	// 获取参数配置
	var port int
	flag.IntVar(&port, "port", Config.Server.Port, "linsten port")
	flag.Parse()

	s := bootstrap.NewServer()
	s.Start(port)
}

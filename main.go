package main

import (
	"GoEchoton/bootstrap"
	"flag"
)

func main() {

	// 获取参数配置
	var port int
	flag.IntVar(&port, "port", 1323, "linsten port")
	flag.Parse()

	s := bootstrap.NewServer()
	s.Start(port)

}

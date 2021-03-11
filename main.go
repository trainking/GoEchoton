package main

import (
	"GoEchoton/pkg/etcd"
	"flag"
	"fmt"
)

func main() {

	// 获取参数配置
	// var port int
	// flag.IntVar(&port, "port", Config.Server.Port, "linsten port")
	// flag.Parse()

	// s := bootstrap.NewServer()
	// s.Start(port)

	var etcdGateway string
	flag.StringVar(&etcdGateway, "etcd", "", "etcd gateway")
	flag.Parse()

	if etcdGateway == "" {
		panic("need a etcd gateway!")
	}

	v, err := etcd.Version(etcdGateway)
	if err != nil {
		fmt.Printf("error:%s\n", err.Error())
	}
	fmt.Println(v)
}

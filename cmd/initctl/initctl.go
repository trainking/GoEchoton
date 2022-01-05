package main

import (
	"GoEchoton/cmd/initctl/options"
	"flag"
	"fmt"
	"strings"
)

var option = flag.String("option", "", "option: conf")
var withCongfigs = flag.String("configs", "", "configs path")
var etcdkey = flag.String("etcdkey", "", "etcd key path")
var etcdGateway = flag.String("etcd", "127.0.0.1:2379", "etcd endpoint to connect to")

func main() {
	flag.Parse()

	var err error
	switch *option {
	case "":
		fmt.Println("no options")
	case "conf":
		err = options.LoadConfigToEtcd(strings.Split(*etcdGateway, ","), *withCongfigs, *etcdkey)
	}
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

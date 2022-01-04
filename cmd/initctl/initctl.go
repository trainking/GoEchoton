package main

import (
	"GoEchoton/pkg/etcdx"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
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
		err = loadConfigToEtcd(strings.Split(*etcdGateway, ","), *withCongfigs, *etcdkey)
	}
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

// loadConfigToEtcd 加载配置到Etcd中
func loadConfigToEtcd(etcdGateway []string, withCongfigs, etcdkey string) error {
	s, err := ioutil.ReadFile(withCongfigs)
	if err != nil {
		return err
	}
	c := etcdx.New(etcdGateway)
	return c.Put(context.TODO(), etcdkey, string(s))
}

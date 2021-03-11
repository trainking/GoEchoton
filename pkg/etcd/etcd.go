package etcd

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// Version 获取etcd版本
func Version(etcdGateway string) (string, error) {
	url := fmt.Sprintf("http://%s/version", etcdGateway)

	fmt.Println(url)
	status, resp, err := fasthttp.Get(nil, url)
	if err != nil {
		return "", err
	}
	if status != fasthttp.StatusOK {
		return "", fmt.Errorf("errcode:%d, body:%v", status, resp)
	}
	return string(resp), nil
}

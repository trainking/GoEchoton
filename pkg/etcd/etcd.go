package etcd

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// EtcdHttpClient http客户端实现
type EtcdHttpClient struct {
	Gateway string
}

// Version 获取etcd版本
func (this *EtcdHttpClient) Version() (string, error) {
	url := fmt.Sprintf("http://%s/version", this.Gateway)

	status, resp, err := fasthttp.Get(nil, url)
	if err != nil {
		return "", err
	}
	if status != fasthttp.StatusOK {
		return "", fmt.Errorf("errcode:%d, body:%v", status, resp)
	}
	return string(resp), nil
}

// Keys 获取所有键值
func (this *EtcdHttpClient) Keys() (string, error) {
	url := fmt.Sprintf("http://%s/v2/keys", this.Gateway)

	status, resp, err := fasthttp.Get(nil, url)
	if err != nil {
		return "", err
	}
	if status != fasthttp.StatusOK {
		return "", fmt.Errorf("errcode:%d, body:%v", status, resp)
	}
	return string(resp), nil
}

// NewHttpClient 创建客户端
func NewHttpClient(etcdGateway string) *EtcdHttpClient {
	return &EtcdHttpClient{Gateway: etcdGateway}
}

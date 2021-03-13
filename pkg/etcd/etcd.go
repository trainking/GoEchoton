package etcd

import (
	"encoding/base64"
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
	url := fmt.Sprintf("http://%s/v3/kv/ragne", this.Gateway)

	status, resp, err := fasthttp.Get(nil, url)
	if err != nil {
		return "", err
	}
	if status != fasthttp.StatusOK {
		return "", fmt.Errorf("errcode:%d, body:%v", status, resp)
	}
	return string(resp), nil
}

// Put put操作
func (this *EtcdHttpClient) Put(key string, value string) error {
	url := fmt.Sprintf("http://%s/v3/kv/put", this.Gateway)

	req := &fasthttp.Request{}
	req.SetRequestURI(url)

	// Etcd 3 使用base64编码
	key_base64 := base64.StdEncoding.EncodeToString([]byte(key))
	value_base64 := base64.StdEncoding.EncodeToString([]byte(value))
	requestBody := []byte(fmt.Sprintf(`{"key":"%s","value":"%s"}`, key_base64, value_base64))
	req.SetBody(requestBody)
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	resp := &fasthttp.Response{}

	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		return err
	}

	b := resp.Body()
	fmt.Println(string(b))
	return nil
}

// Get
func (this *EtcdHttpClient) Get(key string) (string, error) {
	url := fmt.Sprintf("http://%s/v3/kv/range", this.Gateway)

	req := &fasthttp.Request{}
	req.SetRequestURI(url)

	// Etcd 3 使用base64编码
	key_base64 := base64.StdEncoding.EncodeToString([]byte(key))
	requestBody := []byte(fmt.Sprintf(`{"key":"%s"}`, key_base64))
	req.SetBody(requestBody)
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	resp := &fasthttp.Response{}

	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		return "", err
	}

	b := resp.Body()
	// TODO decode response
	return string(b), nil
}

// NewHttpClient 创建客户端
func NewHttpClient(etcdGateway string) *EtcdHttpClient {
	return &EtcdHttpClient{Gateway: etcdGateway}
}

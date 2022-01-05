package options

import (
	"GoEchoton/pkg/etcdx"
	"context"
	"io/ioutil"
)

// loadConfigToEtcd 加载配置到Etcd中
func LoadConfigToEtcd(etcdGateway []string, withCongfigs, etcdkey string) error {
	s, err := ioutil.ReadFile(withCongfigs)
	if err != nil {
		return err
	}
	c := etcdx.New(etcdGateway)
	return c.Put(context.TODO(), etcdkey, string(s))
}

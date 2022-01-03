package apiserver

import (
	"GoEchoton/pkg/etcdx"
	"context"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// LoadConfigFile 从文件中加载配置
func LoadConfigFile(path string, c interface{}) error {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, c)
	return err
}

// LoadConfigEtcdX 从Etcd中加载配置
func LoadConfigEtcdX(path string, client *etcdx.ClientX, c interface{}) error {
	yamlFile, err := client.Get(context.TODO(), path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(yamlFile), c)
	return err
}

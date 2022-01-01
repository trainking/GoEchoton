package apiserver

import (
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

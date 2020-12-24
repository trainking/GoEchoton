package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Yml 配置文件映射结构体
type Yml struct {
	Jwt struct {
		Secret string `yaml:"secret"`
	}
	Mongo struct {
		User   string `yaml:"user"`
		Passwd string `yaml:"passwd"`
		Host   string `yaml:"host"`
		Port   int    `yaml:"port"`
	}
	Server struct {
		Port int `yaml:"port"`
	}
	Redis struct {
		Host   string `yaml:"host"`
		Passwd string `yaml:"passwd"`
		DB     int    `yaml:"db"`
	}
}

// 配置
var Config Yml

func init() {
	path, _ := os.Getwd() // 获取到的是项目根路径 GoEchoton/
	yamlFile, err := ioutil.ReadFile(path + string(os.PathSeparator) + "env.yaml")
	if err != nil {
		panic("Load Env yaml Failed")
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic("Unmarshal: " + err.Error())
	}
}

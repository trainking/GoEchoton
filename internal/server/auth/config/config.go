package config

import "GoEchoton/pkg/logger"

// Config 定义使用配置
type Config struct {
	LoggerConfig logger.Config `yaml:"logger"`
	IpWhiteList  []string      `yaml:"ip_white_list"`
	IpBlackList  []string      `yaml:"ip_black_list"`
}

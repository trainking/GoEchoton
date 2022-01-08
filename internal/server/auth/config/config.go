package config

import "GoEchoton/pkg/logger"

// Config 定义使用配置
type Config struct {
	LoggerConfig logger.Config `yaml:"logger"`
}

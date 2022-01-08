package logger

import (
	"fmt"
	"path/filepath"
	"strings"
)

type Config struct {
	Level   string `json:"level" yaml:"level"`     // 输出的日志级别, 默认debug
	Target  string `json:"target" yaml:"target"`   // 日志标识，一般使用服务名，必须
	ID      string `json:"id" yaml:"id"`           // 与target组成唯一标识
	OutPath string `json:"outPath" yaml:"outPath"` // 日志文件输出路径，默认logs
	MaxSize int    `json:"maxSize" yaml:"maxSize"` // 文件最大大小 MB，默认50M
}

// LogPath 日志文件路径
func (c *Config) LogPath() string {
	// 替换掉可能导致路径问题的字符
	rep := strings.NewReplacer(
		`\`, "",
		"/", "",
		":", "_",
		"*", "_",
		"?", "_",
		"=", "_",
		"<", "_",
		">", "_",
		"|", "_",
	)
	return filepath.Join(c.OutPath, rep.Replace(c.Target), rep.Replace(c.ServiceID()+".log"))
}

// ServiceID 服务ID, 使用Target+ID生成
func (c *Config) ServiceID() string {
	if c.ID == "" {
		return c.Target
	}
	return fmt.Sprintf("%s_%s", c.Target, c.ID)
}

func (c *Config) defaultChange() {
	if c.Level == "" {
		c.Level = "debug"
	}
	if c.Target == "" {
		panic("target is must be")
	}
	if c.OutPath == "" {
		c.OutPath = "logs"
	}
	if c.MaxSize == 0 {
		c.MaxSize = 50
	}
}

package arpcserver

// ServerContext 定义初始化必须的实现
type ServerContext interface {
	// GetHandler 定义获取Handler
	GetHandlers() []Handler
}

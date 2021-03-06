package repository

type Repository interface {

	// Destory 清理释放资源的方法
	Destory() error
}

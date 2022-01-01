package apply

// PageApply 分页请求字段
type PageApply struct {
	Page int64 `json:"page"` // 页码
	Size int64 `json:"size"` // 每页显示数目
}

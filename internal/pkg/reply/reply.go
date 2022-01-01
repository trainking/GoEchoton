package reply

// PageList 分页列表
type PageListReply struct {
	List  interface{} `json:"list"`  // 列表
	Page  int64       `json:"page"`  // 页码
	Size  int64       `json:"size"`  // 每页显示数目
	Total int64       `json:"total"` // 总数
}

// SumPageList 带有小计和总计分页列表
type SumPageListReply struct {
	PageList
	SubTotal interface{} `json:"subtotal"` // 小计
	AllTotal interface{} `json:"alltotal"` // 总计
}

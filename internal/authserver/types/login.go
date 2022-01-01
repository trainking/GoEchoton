package types

type (
	// LoginOneApply 登录第一步请求
	LoginOneApply struct {
		Account string `json:"account"`
	}

	// LoginOneReply 登录第一步返回
	LoginOneReply struct{}
)

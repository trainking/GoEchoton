package model

import "time"

// Hauthorized 验证
type Hauthorized struct {
	Username string
	Tott     uint32    // 不保存token本体，而是保存crc32校验码便于查询
	Date     time.Time // 通过索引设计超时删除
}

const (
	HauthorizedDatabase   = "local"       // 所在数据库
	HauthorizedCollection = "hauthorized" // 集合名
)

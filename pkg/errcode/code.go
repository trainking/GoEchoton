package errcode

//go:generate stringer -type ErrCode
type ErrCode int

const (
	ERR_CODE_OK             ErrCode = 0 // OK
	ERR_CODE_INVALID_PARAMS ErrCode = 1 // 无效参数
	ERR_CODE_TIMEOUT        ErrCode = 2 // 超时
	ERR_CODE_INNER          ErrCode = 3 // 内部错误
)

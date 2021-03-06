package handler

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

const (
	SuccessCode    = iota
	InnerErrorCode // 内部错误
)

// Resp 返回内容结构体
type Resp struct {
	Code int         `json:"code"` // 返回码
	Msg  string      `json:"msg"`  // 消息
	Data interface{} `json:"data"` // 数据
}

// jsoniterResponse 使用jsonIter库
func jsoniterResponse(c echo.Context, v interface{}) error {
	var json = jsoniter.Config{
		EscapeHTML:                    false,
		MarshalFloatWith6Digits:       true, // will lose precession
		ObjectFieldMustBeSimpleString: true, // do not unescape object field
		UseNumber:                     true, // 使用json.Number代替使用flat64表达number
	}.Froze()
	response := c.Response()
	enc := json.NewEncoder(response)
	enc.SetIndent("", "  ")
	header := c.Response().Header()
	if header.Get(echo.HeaderContentType) == "" {
		header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	}
	response.Status = http.StatusOK
	c.SetResponse(response)
	return enc.Encode(v)
}

// Response 成功返回
func Response(c echo.Context, data ...interface{}) error {
	r := Resp{
		Code: SuccessCode,
		Msg:  "ok",
	}
	if len(data) > 0 {
		r.Data = data[0]
	}
	return jsoniterResponse(c, r)
}

// ErrorResponse 错误返回
func ErrorResponse(c echo.Context, msg string, code ...int) error {
	var _code = InnerErrorCode
	if len(code) > 0 {
		_code = code[0]
	}
	r := Resp{
		Code: _code,
		Msg:  msg,
	}
	return jsoniterResponse(c, r)
}

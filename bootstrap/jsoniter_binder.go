package bootstrap

import (
	"net/http"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

// JsoniterBinder 使用jsoniter作为binder
type JsoniterBinder struct{}

// Bind
func (this *JsoniterBinder) Bind(i interface{}, c echo.Context) (err error) {
	req := c.Request()
	ctype := req.Header.Get(echo.HeaderContentType)
	if strings.HasPrefix(ctype, echo.MIMEApplicationJSON) {
		var json = jsoniter.Config{
			EscapeHTML:                    false,
			MarshalFloatWith6Digits:       true, // will lose precession
			ObjectFieldMustBeSimpleString: true, // do not unescape object field
			UseNumber:                     true, // 使用json.Number代替使用flat64表达number
		}.Froze()
		if err = json.NewDecoder(req.Body).Decode(i); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
		}
	}
	// 默认
	db := new(echo.DefaultBinder)
	if err = db.Bind(i, c); err != echo.ErrUnsupportedMediaType {
		return
	}
	return
}

// NewJsoniterBinder
func NewJsoniterBinder() echo.Binder {
	return &JsoniterBinder{}
}

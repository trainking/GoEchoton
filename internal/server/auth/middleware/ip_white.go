package middleare

import (
	"GoEchoton/internal/pkg/ip"
	"net/http"

	"github.com/labstack/echo/v4"
)

// IpWhite 白名单验证器
func IpWhite(ipWhiteList []string) echo.MiddlewareFunc {
	// 加载IP白名单验证器
	ipWhite := ip.NewIpBill(ipWhiteList)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			// 验证
			if !ipWhite.IsDisable() && !ipWhite.Contains(c.RealIP()) {
				return &echo.HTTPError{
					Code: http.StatusNotFound,
				}
			}
			return next(c)
		}
	}
}

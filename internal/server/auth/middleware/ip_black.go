package middleare

import (
	"GoEchoton/internal/pkg/ip"
	"net/http"

	"github.com/labstack/echo/v4"
)

// IpWhite 黑名单验证器
func IpBlack(ipBlackList []string) echo.MiddlewareFunc {
	// 加载IP黑名单验证器
	ipWhite := ip.NewIpBill(ipBlackList)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			// 验证
			if !ipWhite.IsDisable() && ipWhite.Contains(c.RealIP()) {
				return &echo.HTTPError{
					Code: http.StatusNotFound,
				}
			}
			return next(c)
		}
	}
}

package middleare

import (
	"GoEchoton/internal/pkg/ip"
	"net/http"

	"github.com/labstack/echo/v4"
)

func IpWhite() echo.MiddlewareFunc {
	// 加载IP白名单验证器 TODO 配置中获取
	ipWhite := ip.NewIpBill([]string{"127.0.0.1"})
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

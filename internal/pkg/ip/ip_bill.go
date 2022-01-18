package ip

import "net"

type (
	IpBill interface {

		// Contains 是否在白名单中
		Contains(ip string) bool

		// IsDisable 是否未开启
		IsDisable() bool
	}

	defaultIpBill struct {
		whiteList []string // 白名单
		whiteNet  []*net.IPNet
		disable   bool // 白名单是否关闭着
	}
)

// NewIpWhite whitelist必须是IP,或者是RFC4632，和RFC4291的网络号
func NewIpBill(whiteList []string) IpBill {
	var wn []*net.IPNet
	var wi []string
	for _, w := range whiteList {
		_, ipNet, err := net.ParseCIDR(w)
		if err == nil {
			wn = append(wn, ipNet)
		}
		wi = append(wi, w)
	}

	return &defaultIpBill{
		whiteList: wi,
		whiteNet:  wn,
		disable:   len(wi) == 0 && len(wn) == 0,
	}
}

func (d *defaultIpBill) Contains(ip string) bool {
	for _, i := range d.whiteList {
		if ip == i {
			return true
		}
	}
	for _, n := range d.whiteNet {
		if n.Contains(net.ParseIP(ip)) {
			return true
		}
	}
	return false
}

func (d *defaultIpBill) IsDisable() bool {
	return d.disable
}

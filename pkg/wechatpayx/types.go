package wechatpayx

// Config 基本配置
type Config struct {
	MchAppid   string `json:"mch_appid"`   // 申请商户号的appid或商户号绑定的appid
	Mchid      string `json:"mchid"`       // 微信支付分配的商户号
	DeviceInfo string `json:"device_info"` // 微信支付分配的终端设备号, 可无
	SecretKey  string `json:"secret_key"`  // 签名密钥
	// TODO 证书的路径
}

// CompayToUserCoinParams 企业付款到零钱参数
type CompayToUserCoinParams struct {
	PartnerTradNo  string `json:"partner_trade_no"` // 商户订单号，需保持唯一性
	Openid         string `json:"openid"`           // 商户appid下，某用户的openid
	CheckName      string `json:"check_name"`       // NO_CHECK：不校验真实姓名 FORCE_CHECK：强校验真实姓名
	ReUserName     string `json:"re_user_name"`     // 收款用户为FORCE_CHECK有值
	Amount         int64  `json:"amount"`           // 金额， 单位为分
	Desc           string `json:"desc"`             // 付款备注
	SpbillCreateIp string `json:"spbill_create_ip"` // 用户的ip
}

// CompayToUserCoinResult 企业付款到零钱返回结果
type CompayToUserCoinResult struct {
	ReturnCode    string `json:"return_code"`      // SUCCESS/FAIL 这个字段表示通信成功与否
	ReturnMsg     string `json:"return_msg"`       // 通信错误返回信息
	MchAppid      string `json:"mch_appid"`        // 申请商户号的appid或商户号绑定的appid
	Mchid         string `json:"mchid"`            // 微信支付分配的商户号
	DeviceInfo    string `json:"device_info"`      // 微信支付分配的终端设备号, 可无
	NonceStr      string `json:"nonce_str"`        // 随机字符串，不长于32位
	ResultCode    string `json:"result_code"`      // SUCCESS/FAIL
	ErrCode       string `json:"err_code"`         // 错误码信息
	ErrCodeDes    string `json:"err_code_des"`     // 结果信息描述
	PartnerTradNo string `json:"partner_trade_no"` // 商户订单号，需保持唯一性
	PaymentNo     string `json:"payment_no"`       // 付款成功，微信付款单号
	PaymentTime   string `json:"payment_time"`     // 付款成功时间
}

package wechatpayx

import "strconv"

// Config 基本配置
type Config struct {
	MchAppid       string `json:"mch_appid" xml:"mch_appid"`     // 申请商户号的appid或商户号绑定的appid
	Mchid          string `json:"mchid" xml:"mchid"`             // 微信支付分配的商户号
	DeviceInfo     string `json:"device_info" xml:"device_info"` // 微信支付分配的终端设备号, 可无
	SecretKey      string `json:"secret_key"`                    // 签名密钥
	CertPemPath    string `json:"-"`                             // 证书所在路径
	CertKeyPemPath string `json:"-"`                             // 证书密钥所在路径
}

// CompayToUserCoinParams 企业付款到零钱参数
type CompayToUserCoinParams struct {
	Config
	PartnerTradNo  string `json:"partner_trade_no" xml:"partner_trade_no"` // 商户订单号，需保持唯一性
	Openid         string `json:"openid" xml:"openid"`                     // 商户appid下，某用户的openid
	CheckName      string `json:"check_name" xml:"check_name"`             // NO_CHECK：不校验真实姓名 FORCE_CHECK：强校验真实姓名
	ReUserName     string `json:"re_user_name" xml:"re_user_name"`         // 收款用户为FORCE_CHECK有值
	Amount         int64  `json:"amount" xml:"amount"`                     // 金额， 单位为分
	Desc           string `json:"desc" xml:"desc"`                         // 付款备注
	SpbillCreateIp string `json:"spbill_create_ip" xml:"spbill_create_ip"` // 用户的ip
	NonceStr       string `json:"nonce_str" xml:"nonce_str"`               // 随机字符串
	Sign           string `json:"sign" xml:"sign"`                         // 签名
}

// Result 标准返回结果
type Result struct {
	ReturnCode string `json:"return_code" xml:"return_code"`   // SUCCESS/FAIL 这个字段表示通信成功与否
	ReturnMsg  string `json:"return_msg" xml:"return_msg"`     // 通信错误返回信息
	ResultCode string `json:"result_code" xml:"result_code"`   // SUCCESS/FAIL
	ErrCode    string `json:"err_code" xml:"err_code"`         // 错误码信息
	ErrCodeDes string `json:"err_code_des" xml:"err_code_des"` // 结果信息描述
}

// CompayToUserCoinResult 企业付款到零钱返回结果
type CompayToUserCoinResult struct {
	Result
	MchAppid      string `json:"mch_appid" xml:"mch_appid"`               // 申请商户号的appid或商户号绑定的appid
	Mchid         string `json:"mchid" xml:"mchid"`                       // 微信支付分配的商户号
	DeviceInfo    string `json:"device_info" xml:"device_info"`           // 微信支付分配的终端设备号, 可无
	NonceStr      string `json:"nonce_str" xml:"nonce_str"`               // 随机字符串，不长于32位
	PartnerTradNo string `json:"partner_trade_no" xml:"partner_trade_no"` // 商户订单号，需保持唯一性
	PaymentNo     string `json:"payment_no" xml:"payment_no"`             // 付款成功，微信付款单号
	PaymentTime   string `json:"payment_time" xml:"payment_time"`         // 付款成功时间
}

// GetTransferInfoResult 查询付款订单
type GetTransferInfoResult struct {
	Result
	PartnerTradNo string `json:"partner_trade_no" xml:"partner_trade_no"` // 商户订单号，需保持唯一性
	MchAppid      string `json:"appid" xml:"appid"`                       // 申请商户号的appid或商户号绑定的appid
	Mchid         string `json:"mch_id" xml:"mch_id"`                     // 微信支付分配的商户号
	DetailId      string `json:"detail_id" xml:"detail_id"`               // 调用付款API时，微信系统内部产生的单号
	Status        string `json:"status" xml:"status"`                     // SUCCESS:转账成功 FAILED:转账失败 PROCESSING:处理中
	Reason        string `json:"reason" xml:"reason"`                     // 如果失败则有失败原因
	Openid        string `json:"openid" xml:"openid"`                     // 商户appid下，某用户的openid
	TransferName  string `json:"transfer_name" xml:"transfer_name"`       // 收款用户姓名
	PaymentAmount int64  `json:"payment_amount" xml:"payment_amount"`     // 付款金额单位为“分”
	TransferTime  string `json:"transfer_time" xml:"transfer_time"`       // 发起转账的时间
	PaymentTime   string `json:"payment_time" xml:"payment_time"`         // 付款成功时间
	Desc          string `json:"desc" xml:"desc"`                         // 付款备注
}

// getValues 转换类型
func (p *CompayToUserCoinParams) getValues() map[string]string {
	var values = make(map[string]string)

	values["mch_appid"] = p.MchAppid
	values["mchid"] = p.Mchid
	if p.DeviceInfo != "" {
		values["device_info"] = p.DeviceInfo
	}
	values["partner_trade_no"] = p.PartnerTradNo
	values["openid"] = p.Openid
	values["check_name"] = p.CheckName
	if p.CheckName == "FORCE_CHECK" {
		values["re_user_name"] = p.ReUserName
	}
	values["amount"] = strconv.FormatInt(p.Amount, 10)
	values["desc"] = p.Desc
	if p.SpbillCreateIp != "" {
		values["spbill_create_ip"] = p.SpbillCreateIp
	}
	return values
}

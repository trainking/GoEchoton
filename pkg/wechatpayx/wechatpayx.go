package wechatpayx

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"sort"

	"github.com/go-resty/resty/v2"
)

const (
	ReturnCodeSuccess            = "SUCCESS"
	ReturnCodeFail               = "FAIL"
	ErrCodeNoAuTH                = "NO_AUTH"                  // 没有该接口权限
	ErrCodeAmountLimit           = "AMOUNT_LIMIT"             // 金额超额
	ErrCodeParamError            = "PARAM_ERROR"              // 参数错误
	ErrCodeOpenidError           = "OPENID_ERROR"             // Openid错误
	ErrCodeSendFailed            = "SEND_FAILED"              // 付款错误
	ErrCodeNotEnough             = "NOTENOUGH"                // 余额不足
	ErrCodeSystemError           = "SYSTEMERROR"              // 系统繁忙，请稍后再试
	ErrCodeNameMismatch          = "NAME_MISMATCH"            // 收款人姓名校验出错
	ErrCodeSignError             = "SIGN_ERROR"               // 签名错误
	ErrCodeXmlError              = "XML_ERROR"                // xml格式问题
	ErrCodeFatalError            = "FATAL_ERROR"              // 两次请求参数不一致
	ErrCodeFreqLimit             = "FREQ_LIMIT"               // 超过频率限制，请稍后再试
	ErrCodeMoneyLimit            = "MONEY_LIMIT"              // 已经达到今日付款总额上限/已达到付款给此用户额度上限
	ErrCodeCAError               = "CA_ERROR"                 // 商户API证书校验出错
	ErrCodeV2AccountSimpleBan    = "V2_ACCOUNT_SIMPLE_BAN"    // 无法给未实名用户付款
	ErrCodeParamIsNotUtf8        = "PARAM_IS_NOT_UTF8"        // 接口规范要求所有请求参数都必须为utf8编码
	ErrCodeSendnumLimit          = "SENDNUM_LIMIT"            // 该用户今日付款次数超过限制
	ErrCodeRecvAccountNotAllowed = "RECV_ACCOUNT_NOT_ALLOWED" // 收款账户不在收款账户列表
	ErrCodePayChannelNotAllowed  = "PAY_CHANNEL_NOT_ALLOWED"  // 本商户号未配置API发起能力
)

// nonce_str 的元素集
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"

type (
	// Client 发起支付客户端
	Client interface {

		// CompayToUserCoin 企业付款到个人零钱
		CompayToUserCoin(ctx context.Context, params CompayToUserCoinParams) (*CompayToUserCoinResult, error)
	}

	defaultClient struct {
		conf       Config
		httpClient *resty.Client
	}
)

// NewClient 获取一个新Client的唯一途径
// 检查config的参数是否齐全
func NewClient(conf Config) (Client, error) {
	if conf.MchAppid == "" {
		return nil, errors.New("mch_aapid is must be")
	}
	if conf.Mchid == "" {
		return nil, errors.New("mchid is must be")
	}
	if conf.SecretKey == "" {
		return nil, errors.New("secret_key is must be")
	}

	// 设置证书
	cert, err := tls.LoadX509KeyPair(conf.CertPemPath, conf.CertKeyPemPath)
	if err != nil {
		return nil, err
	}
	httpClient := resty.New()
	httpClient.SetCertificates(cert)

	return &defaultClient{conf: conf, httpClient: httpClient}, nil
}

// CompayToUserCoin 企业付款到个人零钱
func (c *defaultClient) CompayToUserCoin(ctx context.Context, params CompayToUserCoinParams) (*CompayToUserCoinResult, error) {
	values := c.conf.getValues()
	for k, v := range params.getValues() {
		values[k] = v
	}
	values["nonce_str"] = c.generateNonceStr()
	sign, err := c.signurate(values, c.conf.SecretKey)
	if err != nil {
		return nil, err
	}
	values["sign"] = sign

	// send to wechatpay
	request := c.httpClient.R()
	request = request.SetHeader("Content-Type", "application/xml")
	request = request.SetBody(values)
	response, err := request.Post("https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers")
	if err != nil {
		return nil, err
	}
	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("CompayToUserCoin send Failed! Code: %d", response.StatusCode())
	}

	var result CompayToUserCoinResult
	if err := xml.Unmarshal(response.Body(), &result); err != nil {
		return nil, err
	}
	// 检查通信是否成功
	if result.ReturnCode == ReturnCodeFail {
		return nil, fmt.Errorf("Wechatpay Error: %s", result.ReturnMsg)
	}
	// 检查支付是否成功
	if result.ResultCode == ReturnCodeFail {
		return nil, fmt.Errorf("Wechatpay Error: %s", result.ReturnMsg)
	}
	return &result, nil
}

// generateNonceStr 随机产生32位随机字符串
func (c *defaultClient) generateNonceStr() string {
	n := rand.Intn(32)
	if n < 10 {
		n = 10
	}
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

// signurate 求出签名
func (c *defaultClient) signurate(values map[string]string, secret_key string) (string, error) {
	var keys []string
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buff bytes.Buffer
	for _, k := range keys {
		buff.WriteString(fmt.Sprintf("%s=%s&", k, values[k]))
	}
	buff.WriteString(fmt.Sprintf("key=%s", secret_key))
	signStr := buff.String()
	md5W := md5.New()
	if _, err := io.WriteString(md5W, signStr); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", md5W.Sum(nil)), nil
}

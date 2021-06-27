package wechatpayx

import "testing"

func TestGenerateNonceStr(t *testing.T) {
	clinet := defaultClient{}
	for i := 0; i < 3; i++ {
		nonce_str := clinet.generateNonceStr()
		t.Log(nonce_str)
	}
}

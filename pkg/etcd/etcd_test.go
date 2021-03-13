package etcd

import "testing"

func TestEtcdHtppClientVersion(t *testing.T) {
	ht := NewHttpClient("192.168.33.10:2379")
	v, err := ht.Version()
	if err != nil {
		t.Error(err)
	}
	t.Log(v)
}

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

func TestEtcdHtppClientPut(t *testing.T) {
	ht := NewHttpClient("192.168.33.10:2379")
	err := ht.Put("hello", "world")
	if err != nil {
		t.Error(err)
	}
}

func TestEtcdHtppClientGet(t *testing.T) {
	ht := NewHttpClient("192.168.33.10:2379")
	value, err := ht.Get("hello")
	if err != nil {
		t.Error(err)
	}
	t.Log(value)
}

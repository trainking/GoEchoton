package arpcx

import (
	"fmt"
	"testing"

	"github.com/lesismal/arpc"
)

type FakeClientPool struct{}

func NewFakeClientPool() ClientPool {
	return &FakeClientPool{}
}

func (f *FakeClientPool) Create(listenAddr string, size int) (*arpc.ClientPool, error) {
	return &arpc.ClientPool{}, nil
}

func (f *FakeClientPool) Choose(p *arpc.ClientPool) *arpc.Client {
	return &arpc.Client{}
}

func TestAddClientPool(t *testing.T) {
	pools := make(map[string]*arpc.ClientPool)
	c := &Client{pools: pools, cp: NewFakeClientPool()}

	if err := c.AddNode("127.0.0.1:80001", 2); err != nil {
		t.Error(err)
	}

	if c.length != 1 {
		t.Errorf("expected 1 client pool")
	}

	if len(c.addrs) != len(c.pools) {
		t.Error("addrs not equal pools")
	}
}

func BenchmarkAddClientPool(b *testing.B) {
	pools := make(map[string]*arpc.ClientPool)
	c := &Client{pools: pools, cp: NewFakeClientPool()}

	for i := 0; i < b.N; i++ {
		if err := c.AddNode(fmt.Sprintf("127.0.0.1:8000%d", i), 2); err != nil {
			b.Error(err)
		}
	}
}

func TestDeleteClientPool(t *testing.T) {
	pools := make(map[string]*arpc.ClientPool)
	c := &Client{pools: pools, cp: NewFakeClientPool()}

	if err := c.AddNode("127.0.0.1:80001", 2); err != nil {
		t.Error(err)
	}

	if err := c.AddNode("127.0.0.1:80002", 2); err != nil {
		t.Error(err)
	}

	if err := c.AddNode("127.0.0.1:80003", 2); err != nil {
		t.Error(err)
	}

	c.DeleteNode("127.0.0.1:80001")

	if c.length != 2 {
		t.Errorf("expected 2 client pools")
	}

	if len(c.addrs) != len(c.pools) {
		t.Errorf("addrs not equal pools")
	}

	for _, s := range c.addrs {
		if _, ok := c.pools[s]; !ok {
			t.Errorf("delete error key")
		}
	}

}

func BenchmarkDeleteClientPool(b *testing.B) {
	pools := make(map[string]*arpc.ClientPool)
	c := &Client{pools: pools, cp: NewFakeClientPool()}

	var keys []string
	for j := 0; j < 1000; j++ {
		k := fmt.Sprintf("172.0.0.1:%d", j)
		keys = append(keys, k)
		if err := c.AddNode(k, 2); err != nil {
			b.Error(err)
		}
	}

	for i := 0; i < b.N; i++ {
		c.DeleteNode(keys[i%1000])
	}

}

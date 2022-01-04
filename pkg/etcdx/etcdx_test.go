package etcdx

import (
	"context"
	"testing"
)

func TestPut(t *testing.T) {
	c := New([]string{"127.0.0.1:2379"})

	if err := c.Put(context.TODO(), "hello", "world"); err != nil {
		t.Error(err)
	}
}

func TestGet(t *testing.T) {
	c := New([]string{"127.0.0.1:2379"})

	if s, err := c.Get(context.TODO(), "hello"); err != nil {
		t.Error(err)
	} else {
		if s != "world" {
			t.Errorf("got %v, want %v", s, "world")
		}
	}

	if _, err := c.Get(context.TODO(), "hello1"); err != nil {
		t.Error(err)
	}

}

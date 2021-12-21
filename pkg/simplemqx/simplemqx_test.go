package simplemqx

import (
	"fmt"
	"testing"
)

func TestSend(t *testing.T) {
	q := New("amqp://guest:guest@localhost:5672//test", "q_order_payment")
	if err := q.Connect(); err != nil {
		t.Error(err)
	}

	if err := q.Push([]byte("hello,world!")); err != nil {
		t.Error(err)
	}
}

func TestCustme(t *testing.T) {
	q := New("amqp://guest:guest@localhost:5672//test", "q_order_payment")
	if err := q.Connect(); err != nil {
		t.Error(err)
	}

	if err := q.Cousume(func(b []byte) error {
		fmt.Println(string(b))
		return nil
	}); err != nil {
		t.Error(err)
	}
}

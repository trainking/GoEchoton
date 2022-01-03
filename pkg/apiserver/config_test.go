package apiserver

import (
	"GoEchoton/pkg/etcdx"
	"testing"
)

func TestLoadConfigEtcdX(t *testing.T) {
	path := "/authserver/config"
	clientx := etcdx.New([]string{"127.0.0.1:2379"})

	var c struct {
		Port    int64  `yaml:"port"`
		Address string `yaml:"address"`
	}

	if err := LoadConfigEtcdX(path, clientx, &c); err != nil {
		t.Error(err)
	}

	if c.Port != 2344 {
		t.Errorf("Expected 2344, got %v", c.Port)
	}
	t.Logf("%+v", c)
}

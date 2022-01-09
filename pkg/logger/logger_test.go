package logger

import "testing"

func TestInfo(t *testing.T) {
	l := New(Config{
		Level:   "info",
		Target:  "test",
		ID:      "127.0.0.1:8009",
		MaxSize: 50,
	})
	// l.Debug("3232")
	l.Info("hello")
	l.Infof("httllo %s", "dafsd")
	l.Info("ddd")
}

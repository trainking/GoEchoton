package sqlt

import (
	"testing"
	"time"
)

func TestInsertSQL(t *testing.T) {
	g := &InsertSQL{
		Table: "hello",
	}
	g.Add("aa", 111)
	g.Add("pp", "cc")
	g.Add("teime", time.Now())
	t.Fatalf("Data: %v", g.Data)
}

package sqlt

import (
	"testing"
	"time"
)

func TestInsertSQLAdd(t *testing.T) {
	g := &InsertSQL{
		Table: "hello",
	}
	g.Add("aa", 111)
	g.Add("pp", "cc")
	g.Add("teime", time.Now())
	if g.Sql() != "INSERT INTO hello (aa,pp,teime) VALUES (?,?,?)" {
		t.Error(g.Sql())
	}
	if g.Data[0] != 111 {
		t.Error("error data")
	}
}

func TestInsertSQLAddByStruct(t *testing.T) {
	g := &InsertSQL{
		Table: "hello",
	}
	var params = &struct {
		Aa    int       `db:"aa"`
		Pp    string    `db:"pp"`
		Teime time.Time `db:"teime"`
	}{111, "cc", time.Now()}
	g.AddByStruct(params)
	if g.Sql() != "INSERT INTO hello (aa,pp,teime) VALUES (?,?,?)" {
		t.Error(g.Sql())
	}
	if g.Data[0] != 111 {
		t.Error("error data")
	}
}

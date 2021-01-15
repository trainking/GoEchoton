package sqlt

import (
	"bytes"
	"fmt"
)

// InsertSQL 插入语句
type InsertSQL struct {
	Table string        // 表名
	Data  []interface{} // 插入数据
	cbuff bytes.Buffer  // 字段名
	vbuff bytes.Buffer  // 序号
}

// Add 增加数据
func (this *InsertSQL) Add(col string, v interface{}) *InsertSQL {
	this.Data = append(this.Data, v)
	if this.cbuff.Len() != 0 {
		this.cbuff.WriteByte(',')
	}
	this.cbuff.WriteString(col)

	if this.vbuff.Len() != 0 {
		this.vbuff.WriteByte(',')
	}
	this.vbuff.WriteByte('?')

	return this
}

// AddMany 批量增加
func (this *InsertSQL) AddMany(cols map[string]interface{}) *InsertSQL {
	for k, v := range cols {
		this.Add(k, v)
	}
	return this
}

// Sql 获取sql语句
func (this *InsertSQL) Sql() string {
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", this.Table, this.cbuff.String(), this.vbuff.String())
}

// T 返回语句和数据
func (this *InsertSQL) T() (string, []interface{}) {
	return this.Sql(), this.Data
}

package sqlt

import (
	"fmt"
	"reflect"
	"strings"
)

// InsertSQL 插入语句
type InsertSQL struct {
	Table string          // 表名
	Data  []interface{}   // 插入数据
	cbuff strings.Builder // 字段名
	vbuff strings.Builder // 序号
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

// AddByStruct 从结构体增加
func (this *InsertSQL) AddByStruct(col interface{}) *InsertSQL {
	v := reflect.ValueOf(col)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		panic("AddByStruct need a struct")
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		if k := f.Tag.Get("db"); k != "" {
			this.Add(k, v.Field(i).Interface())
		}
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

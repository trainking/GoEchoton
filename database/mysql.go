package database

import (
	. "GoEchoton/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Mysql mysql结构体
type Mysql struct {
	db *sql.DB
}

// NewMysql 新建Mysql连接
func NewMysql() (*Mysql, error) {
	conn := fmt.Sprintf(
		"%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=true",
		Config.Mysql.User,
		Config.Mysql.Passwd,
		Config.Mysql.Network,
		Config.Mysql.Host,
		Config.Mysql.Port,
		Config.Mysql.Database,
	)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}
	return &Mysql{db: db}, nil
}

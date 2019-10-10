package db

import (
	"database/sql"
	"fmt"
	"kboard/config"
	"kboard/utils/exception"
	"kboard/utils"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	User         string
	Passwd       string
	Host         string
	Port         string
	Dbname       string
	Charset      string
	MaxOpenConns int
	Timeout      int
}

func NewMysql(config *config.Config) *Mysql {
	return &Mysql{
		User:         config.Data.Mysql.Username,
		Passwd:       config.Data.Mysql.Password,
		Host:         config.Data.Mysql.Host,
		Port:         utils.ToString(config.Data.Mysql.Port),
		Dbname:       config.Data.Mysql.Dbname,
		Charset:      config.Data.Mysql.Charset,
		MaxOpenConns: config.Data.Mysql.MaxOpenConns,
		Timeout:      60,
	}
}

func (m *Mysql) Init() {
	// 检查连接参数
	if m.User == "" || m.Passwd == "" || m.Host == "" || m.Dbname == "" {
		exception.CheckError(exception.NewError("mysql connection parameters is error"), -1)
	}
	if m.Charset == "" {
		m.Charset = "utf8"
	}
	if m.Port == "" {
		m.Port = "3306"
	}
	mysqlSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&timeout=10s",
		m.User,
		m.Passwd,
		m.Host,
		m.Port,
		m.Dbname,
		m.Charset)

	var err error
	DbConn, err = sql.Open("mysql", mysqlSource)
	exception.CheckError(err, -1)
	DbConn.SetMaxIdleConns(2)
	DbConn.SetMaxOpenConns(m.MaxOpenConns)
	DbConn.SetConnMaxLifetime(time.Duration(m.Timeout) * time.Second)

	log.Println("init mysql")
}

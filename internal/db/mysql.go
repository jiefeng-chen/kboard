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

func NewMysql(config config.IConfig) *Mysql {
	configData := config.GetConfigData()
	return &Mysql{
		User:         configData.Mysql.Username,
		Passwd:       configData.Mysql.Password,
		Host:         configData.Mysql.Host,
		Port:         utils.ToString(configData.Mysql.Port),
		Dbname:       configData.Mysql.Dbname,
		Charset:      configData.Mysql.Charset,
		MaxOpenConns: configData.Mysql.MaxOpenConns,
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

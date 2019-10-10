package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var (
	DbConn *sql.DB
)

type SQLite struct {
	Db     *sql.DB
	DbPath string
}

func NewSQLite() *SQLite {
	SQLiteDb := &SQLite{}

	return SQLiteDb
}

func (s *SQLite) Init(dbPath string) {
	if dbPath == "" {
		panic("sqlite db path error")
	}

	var err error
	DbConn, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}

	s.Db = DbConn
	s.DbPath = dbPath

	log.Println("init sqlite...")
}
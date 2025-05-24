package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(dsn string) error {
    var err error
    DB, err = sql.Open("postgres", dsn)
    return err
}
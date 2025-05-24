package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"strings"
	"os"

)

func SetupDBConnection() (*sql.DB, error) {
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")

	log.Debug("mysql", DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME)

	db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME)
	if err != nil {
	    log.Error("Failed to connect to database")
	}

	return db, nil
}


func Insert(db *sql.DB,table string, columns []string, values []string) {
	log.Debug("insert into " + table + " (" + strings.Join(columns, ", ") + ") values(" + strings.Join(values, ", ") + ")")

	stmtIns, err := db.Prepare("insert into " + table + " (" + strings.Join(columns, ", ") + ") values(" + strings.Join(values, ", ") + ")")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec()
	if err != nil {
		panic(err.Error())
	}
}
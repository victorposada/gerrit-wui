package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strings"
)

func maain() {
	fmt.Println("Hello, world!")
	createDatabase("foo")
	insert("foo", "foo", []string{"id", "name"}, []string{"1", "'bar'"})
}

func createDatabase(database_name string) *sql.DB {
	os.Remove("./" + database_name +".db")

	db, err := sql.Open("sqlite3", "./" + database_name +".db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}
	return db
}

func inserat(database_name string, table string, columns []string, values []string) {

	db, err := sql.Open("sqlite3", "./" + database_name +".db")
	if err != nil {
		log.Fatal(err)
	}

	transaction, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("insert into " + table + " (" + strings.Join(columns, ", ") + ") values(" + strings.Join(values, ", ") + ")")
	statement, err := transaction.Prepare("insert into " + table + " (" + strings.Join(columns, ", ") + ") values(" + strings.Join(values, ", ") + ")")
	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()

	err = transaction.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
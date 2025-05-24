package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	// "os"
	"github.com/joho/godotenv"
	"log"
	"strings"
	"net/http"
	"io"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// DB_HOST := os.Getenv("DB_HOST")
	// DB_PORT := os.Getenv("DB_PORT")
	// DB_USER := os.Getenv("DB_USER")
	// DB_PASS := os.Getenv("DB_PASS")
	// DB_NAME := os.Getenv("DB_NAME")

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	http.ListenAndServe(":3333", nil)

	// db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME)
	// insert(db, "users", []string{"id", "name"}, []string{"2", "'bar'"})

}

func insert(db *sql.DB,table string, columns []string, values []string) {
	fmt.Println("insert into " + table + " (" + strings.Join(columns, ", ") + ") values(" + strings.Join(values, ", ") + ")")

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

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
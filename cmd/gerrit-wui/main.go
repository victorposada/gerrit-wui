package main

import (
    "log"
	"github.com/joho/godotenv"
    "github.com/victorposada/gerrit-wui/internal/server"
    // "github.com/victorposada/gerrit-wui/internal/database"
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

	// err := database.InitDB("user:password@tcp(localhost:3306)/dbname")
	// if err != nil {
	//     log.Fatalf("Failed to connect to database: %v", err)
	// }
	// db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME)
	// insert(db, "users", []string{"id", "name"}, []string{"2", "'bar'"})

	// Start the server
	server.StartServer()
}
package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/joho/godotenv"
    "github.com/victorposada/gerrit-wui/internal/web"
	"github.com/victorposada/gerrit-wui/internal/gerrit"
    //"github.com/victorposada/gerrit-wui/internal/db"
	"os"
	//"fmt"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file")
	}

	switch os.Getenv("LOG_LEVEL") {
	case "debug":
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}

	// repos, _ := gerrit.GetProjects()
	gerrit.GetChanges()

	//print(repos)

	// database, _:= db.SetupDBConnection()

	// db.Insert(database, "users", []string{"id", "name"}, []string{"3", "'bar'"})


	// err := database.InitDB("user:password@tcp(localhost:3306)/dbname")
	// if err != nil {
	//     log.Fatalf("Failed to connect to database: %v", err)
	// }
	// db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME)
	// insert(db, "users", []string{"id", "name"}, []string{"2", "'bar'"})

	// Start the server
	web.StartServer()
}
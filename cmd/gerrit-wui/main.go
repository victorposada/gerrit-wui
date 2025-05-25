package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/joho/godotenv"
    "github.com/victorposada/gerrit-wui/internal/web"
	//"github.com/victorposada/gerrit-wui/internal/gerrit"
	"os"

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
	// changes := gerrit.GetChanges()

	//print(repos)
	log.Debug("Pre web-server")
	web.StartServer()
}
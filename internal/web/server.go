package web

import (
    "net/http"
	log "github.com/sirupsen/logrus"
	"os"
)

func StartServer() {
	SERVER_PORT := os.Getenv("SERVER_PORT")
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/info", getInfo)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Info("failed to fetch URL: http://localhost:" + SERVER_PORT + "")
	// Start the server on port 3333
	if err := http.ListenAndServe(":"+ SERVER_PORT, nil); err != nil {
		panic(err)
	}
}
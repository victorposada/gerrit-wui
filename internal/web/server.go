package web

import (
    "net/http"
	log "github.com/sirupsen/logrus"
)

func StartServer() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	log.Debug("failed to fetch URL: http://example.com")
	// Start the server on port 3333
	if err := http.ListenAndServe(":3333", nil); err != nil {
		panic(err)
	}
}
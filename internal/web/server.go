package web

import (
    "net/http"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func StartServer() {
	SERVER_PORT := os.Getenv("SERVER_PORT")
	log.Info("Server up on localhost:" + SERVER_PORT)
	boards := os.Getenv("BOARDS")
	for _, board := range strings.Split(boards, ",") {
		http.HandleFunc("/" + board, getBoard(board))
	}
	http.HandleFunc("/", getBoards)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	if err := http.ListenAndServe(":"+ SERVER_PORT, nil); err != nil {
		panic(err)
	}

}
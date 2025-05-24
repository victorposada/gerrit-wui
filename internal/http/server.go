package http

import (
    "net/http"
)

func startWebServer() {

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	http.ListenAndServe(":3333", nil)

}
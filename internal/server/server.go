package server

import (
    "net/http"
)

func StartServer() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	// Start the server on port 3333
	if err := http.ListenAndServe(":3333", nil); err != nil {
		panic(err)
	}
}
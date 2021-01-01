package main

import (
	"net/http"
)

func main() {
	// 何もしない
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}
	server.ListenAndServe()
}

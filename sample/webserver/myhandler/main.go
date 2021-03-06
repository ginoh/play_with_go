package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

// ServeHTTPを実装しているものが Handler
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &handler,
	}
	server.ListenAndServe()
}

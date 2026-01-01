package main

import (
	"fmt"
	"log"
	"net/http"
)

type MyHandler struct {
	Message string
}

// ServeHTTPを実装しているものが Handler
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", h.Message)
}

func main() {
	handler := MyHandler{
		Message: "Hello World!",
	}
	addr := "localhost:8080"
	log.Println("Starting server on http://" + addr)
	server := http.Server{
		Addr:    addr,
		Handler: &handler,
	}
	log.Fatal(server.ListenAndServe())
}

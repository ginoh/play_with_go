package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// マルチプレクサを生成して、ハンドラとして利用する
	mux := http.NewServeMux()

	mux.HandleFunc("/", hello)

	addr := "localhost:8080"
	log.Println("Starting server on http://" + addr)
	server := &http.Server{
		Addr:    addr,
		Handler: mux, // マルチプレクサは Handler でもある
	}
	log.Fatal(server.ListenAndServe())
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello MultiPlexer")
}

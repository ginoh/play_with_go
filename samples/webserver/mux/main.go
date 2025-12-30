package main

import (
	"fmt"
	"net/http"
)

func main() {
	// multi plexer を生成して、ハンドラとして利用する
	mux := http.NewServeMux()

	mux.HandleFunc("/", hello)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux, // マルチプレクサは Handler でもある
	}
	server.ListenAndServe()
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello MultiPlexer")
}

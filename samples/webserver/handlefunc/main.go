package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	hogeHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Hoge!")
	}
	fooHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Foo!")
	}
	http.HandleFunc("/hoge", hogeHandler)
	http.HandleFunc("/foo", fooHandler)

	addr := "localhost:8080"
	log.Println("Starting server on http://" + addr)
	// nil 指定で defaultのマルチプレクサを利用
	log.Fatal(http.ListenAndServe(addr, nil))
}

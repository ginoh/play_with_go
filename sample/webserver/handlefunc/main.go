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
		fmt.Fprintf(w, "Root Foo!")
	}
	http.HandleFunc("/hoge", hogeHandler)
	http.HandleFunc("/foo", fooHandler)

	log.Println("Startting Server")
	// nil 指定で defaultのマルチプレクサを利用
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

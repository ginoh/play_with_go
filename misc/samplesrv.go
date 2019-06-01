package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	}
	rootHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Root Handler")
	}
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Println("Startting Server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

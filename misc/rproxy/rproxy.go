package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	director := func(r *http.Request) {
		r.URL.Scheme = "http"
		r.URL.Host = ":5000"
	}
	rp := &httputil.ReverseProxy{Director: director}

	log.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":8080", rp))
}

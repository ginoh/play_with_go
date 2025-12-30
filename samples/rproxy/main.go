package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	director := func(r *http.Request) {
		r.URL.Scheme = "http"
		r.URL.Host = "localhost:5000"
	}
	rp := &httputil.ReverseProxy{Director: director}

	statusHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", http.StatusText(http.StatusOK))
	}

	// I can add processing before proxy processing
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rp.ServeHTTP(w, r)
	})

	http.HandleFunc("/status", statusHandler)

	log.Println("Starting Server")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

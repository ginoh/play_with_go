package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil, // default のマルチプレクサの利用
	}
	server.ListenAndServe()
	//server.ListenAndServeTLS("cert.pem", "key.pem") // tls
}

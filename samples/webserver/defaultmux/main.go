package main

import (
	"log"
	"net/http"
)

func main() {
	addr := "localhost:8080"
	log.Println("Starting server on http://" + addr)
	server := http.Server{
		Addr:    addr,
		Handler: nil, // default のマルチプレクサの利用
	}
	log.Fatal(server.ListenAndServe())
	//server.ListenAndServeTLS("cert.pem", "key.pem") // tls
	// 例: openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -days 365
}

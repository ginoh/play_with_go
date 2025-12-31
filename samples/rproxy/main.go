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
		// 上流の Host に合わせてヘッダも更新する
		r.Host = r.URL.Host
	}
	// 失敗時の挙動を明示し、ログとレスポンスを分かりやすくする
	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("proxy error: %v", err)
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
	}
	rp := &httputil.ReverseProxy{
		Director:     director,
		ErrorHandler: errorHandler,
	}

	statusHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", http.StatusText(http.StatusOK))
	}

	// DefaultServeMux は最長一致でマッチする
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rp.ServeHTTP(w, r)
	})

	http.HandleFunc("/status", statusHandler)

	log.Println("Starting Server")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

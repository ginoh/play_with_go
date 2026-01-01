package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// handler をラップして前処理を挟む（簡易ミドルウェア）
func logging(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 実体の関数名を取得してログに出す
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler Function called - " + name)
		h(w, r)
	}
}
func main() {
	// handler を指定しないので default のマルチプレクサの利用
	addr := "localhost:8080"
	log.Println("Starting server on http://" + addr)
	server := http.Server{
		Addr: addr,
	}
	http.HandleFunc("/hello", logging(hello))
	log.Fatal(server.ListenAndServe())
}

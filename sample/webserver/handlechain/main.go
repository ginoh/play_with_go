package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler Function called - " + name)
		h(w, r)
	}
}
func main() {
	// handler を指定しないので default のマルチプレクサの利用
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/hello", log(hello))
	server.ListenAndServe()
}

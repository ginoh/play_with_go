package main

import (
	"fmt"
	"log"
	"net/http"
)

type HogeHandler struct{}

func (h *HogeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Hoge!")
}

type FooHandler struct{}

func (h *FooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Foo!")
}

func main() {
	hoge := HogeHandler{}
	foo := FooHandler{}

	addr := "localhost:8080"
	log.Println("Starting server on http://" + addr)
	server := http.Server{
		Addr: addr,
	}

	// デフォルトマルチプレクサにハンドラを設定
	http.Handle("/hoge", &hoge)
	http.Handle("/foo", &foo)

	log.Fatal(server.ListenAndServe())
}

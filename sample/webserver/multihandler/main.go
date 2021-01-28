package main

import (
	"fmt"
	"net/http"
)

type HogeHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Hoge!")
}

type FooHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Foo!")
}

func main() {
	hoge := HogeHandler{}
	foo := FooHandler{}

	server := http.Server{
		Addr: "localhost:8080",
	}

	// デフォルトマルチプレクサにハンドラを設定
	http.Handle("/hoge", &hoge)
	http.Handle("/foo", &foo)

	server.ListenAndServe()
}

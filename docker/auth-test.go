package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	director := func(r *http.Request) {
		r.URL.Scheme = "https"
		r.URL.Host = "registry-1.docker.io"
		r.URL.Path = "/v2/hoge/foo/manifests/latest"
		r.Host = r.URL.Host // Host Header
	}
	modifier := func(res *http.Response) error {
		fmt.Println(res.Request.URL)
		fmt.Println(res.StatusCode)

		dump, _ := httputil.DumpRequest(res.Request, true)
		fmt.Printf("%q", dump)

		return nil
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		fmt.Printf("%#v\n", err)
	}
	rp := &httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifier,
		ErrorHandler:   errorHandler,
	}

	statusHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", http.StatusText(http.StatusOK))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rp.ServeHTTP(w, r)
	})

	http.HandleFunc("/status", statusHandler)

	log.Println("Starting Server")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

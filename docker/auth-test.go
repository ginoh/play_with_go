package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	// m := challenge.NewSimpleManager()

	// resp, err := http.Get("https://registry-1.docker.io/v2/hoge/foo/manifests/latest")
	// // resp, err := http.Get("https://registry-1.docker.io/v2/")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// if resp.StatusCode != http.StatusUnauthorized {
	// 	fmt.Println("Not Unauthorized Response")
	// }
	// m.AddResponse(resp)

	// ep, err := url.Parse("https://registry-1.docker.io:443/v2/hoge/foo/manifests/latest")
	// // ep, err := url.Parse("https://registry-1.docker.io:443/v2/")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// ch, err := m.GetChallenges(*ep)
	// for _, c := range ch {
	// 	fmt.Printf("%#v", c.Parameters)
	// }

	director := func(r *http.Request) {
		r.URL.Scheme = "https"
		r.URL.Host = "registry-1.docker.io:443"
		r.URL.Path = "/v2/hoge/foo/manifests/latest"
	}
	rp := &httputil.ReverseProxy{Director: director}

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

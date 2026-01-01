package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

func main() {
	f1, err := os.OpenFile("tokyo.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()
	f2, err := os.OpenFile("nagoya.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	// 1回送信の用途なので close は使わない
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	// 実行前に有効なURLへ差し替える（JSONが返る想定）
	// 東京天気
	url1 := "https://example.com/tokyo.json"
	// 名古屋天気
	url2 := "https://example.com/nagoya.json"

	wg := &sync.WaitGroup{}

	wg.Add(2)
	go getWeb(url1, ch1, wg)
	go getWeb(url2, ch2, wg)
	wg.Wait()

	fmt.Println("all goroutine Done")
	if _, err := f1.WriteString(<-ch1); err != nil {
		log.Fatal(err)
	}
	if _, err := f2.WriteString(<-ch2); err != nil {
		log.Fatal(err)
	}
}

func getWeb(url string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		ch <- fmt.Sprintf("http Get error: %v", err)
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		ch <- fmt.Sprintf("Read error: %v", err)
		return
	}
	ch <- string(body)
}

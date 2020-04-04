package main

import (
	"fmt"
	"io/ioutil"
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
	f2, err := os.OpenFile("nagoyo.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	ch1 := make(chan string, 2)
	ch2 := make(chan string, 2)

	// 東京天気
	url1 := "http://weather.livedoor.com/forecast/webservice/json/v1?city=130010"
	// 名古屋天気
	url2 := "http://weather.livedoor.com/forecast/webservice/json/v1?city=230010"

	wg := new(sync.WaitGroup)

	wg.Add(2)
	go getWeb(url1, ch1, wg)
	go getWeb(url2, ch2, wg)
	wg.Wait()

	fmt.Println("all goroutine Done")
	f1.WriteString(<-ch1)
	f2.WriteString(<-ch2)
}

func getWeb(url string, ch chan string, wg *sync.WaitGroup) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		ch <- "http Get error"
		close(ch)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		ch <- "Read error"
		close(ch)
		return
	}
	ch <- string(body)
	wg.Done()
}

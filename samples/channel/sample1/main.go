package main

import (
	"fmt"
)

func main() {

	// channel 作成
	ch := make(chan int)

	// channel は goroutine 間のデータ共有用
	go receiver(ch)

	i := 0
	for i < 100 {
		ch <- i
		i++
	}
}

func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

package main

import "fmt"

func main() {

	ch := make(chan int)

	// channelはgoroutine間のデータ共有用
	go receiver(ch)

	i := 0
	for i < 10000 {
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

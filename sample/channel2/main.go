package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 20)

	// channelはgoroutine間のデータ共有用
	go receiver("1st", ch)
	go receiver("2nd", ch)
	go receiver("3rd", ch)

	i := 0
	for i < 100 {
		ch <- i
		i++
	}
	close(ch)

	time.Sleep(3 * time.Second)
}

func receiver(name string, ch <-chan int) {
	for {
		// バッファ空 && closeの場合は ok == false
		i, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println(name, i)
	}
	// // rangeも利用できるが、closeのタイミングがわからない
	// for i := range ch {
	// 	fmt.Println(name, i)
	// }
	fmt.Println(name + " is done.")
}

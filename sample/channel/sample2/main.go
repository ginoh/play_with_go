package main

import (
	"fmt"
	"time"
)

func main() {

	// channel 作成
	ch := make(chan int, 20)

	// channel は goroutine 間のデータ共有用
	// 処理する receiver を 3つ用意
	go receiver("rv1", ch)
	go receiver("rv2", ch)
	go receiver("rv3", ch)

	i := 0
	for i < 100 {
		ch <- i
		i++
	}
	close(ch)

	// 処理が終わる余裕のために 3sec 待機
	time.Sleep(3 * time.Second)
}

func receiver(name string, ch <-chan int) {
	for {
		// バッファ空 && close の場合は ok == false
		i, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println(name, i)
	}
	// // rangeも利用できるが、close のタイミングがわからない
	// for i := range ch {
	// 	fmt.Println(name, i)
	// }
	fmt.Println(name + " is done.")
}

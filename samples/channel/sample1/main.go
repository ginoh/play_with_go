package main

import (
	"fmt"
	"sync"
)

func main() {
	// channel 作成
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)
	go receiver(ch, &wg)

	for i := 0; i < 100; i++ {
		ch <- i
	}
	// 送信完了を通知して受信側のループを終了させる
	close(ch)

	// 受信側の処理完了を待つ
	wg.Wait()
}

func receiver(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		fmt.Println(i)
	}
}

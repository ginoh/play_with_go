package main

import (
	"fmt"
	"sync"
)

func main() {

	// channel 作成
	ch := make(chan int, 20)

	// channel は goroutine 間のデータ共有用
	// 処理する receiver を 3つ用意
	var wg sync.WaitGroup
	wg.Add(3)
	go receiver("rv1", ch, &wg)
	go receiver("rv2", ch, &wg)
	go receiver("rv3", ch, &wg)

	for i := 0; i < 100; i++ {
		ch <- i
	}
	// 送信完了を通知して受信側のループを終了させる
	close(ch)

	// 受信側の処理完了を待つ
	wg.Wait()
}

func receiver(name string, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// range ならチャネルがクローズされるまで受信し続ける
	for i := range ch {
		fmt.Println(name, i)
	}
	// rangeも利用できる、チャネルがクローズするとループから抜ける
	// for i := range ch {
	// 	fmt.Println(name, i)
	// }
	fmt.Println(name + " is done.")
}

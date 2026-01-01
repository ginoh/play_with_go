package main

import (
	"fmt"
	"time"
)

func main() {
	// 同一ゴルーチンで送受信するため、バッファがないと送信で詰まる
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	ch1 <- 1
	ch2 <- 2

	// 複数のcase節が成立する場合はランダムに選択
	select {
	case <-ch1:
		fmt.Println("ch1から受信")
	case <-ch2:
		fmt.Println("ch2から受信")
	case ch3 <- 3:
		// 送信も select の候補にできる
		fmt.Println("ch3へ送信")
	case <-time.After(1 * time.Second):
		// ch1/ch2 が即時に準備できるため、この例では基本的に到達しない
		fmt.Println("timeout")
	default:
		// どのcaseも準備できていないときは即座にここに来る
		// この例では到達しない（ch1/ch2が準備できている）
		fmt.Println("default")
	}

	// 送信された場合だけ受信する
	select {
	case v := <-ch3:
		fmt.Printf("ch3から受信: %d\n", v)
	default:
		fmt.Println("ch3へ送信されていない")
	}
}

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// 4 sec 時間経過で cancel が実行
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	// WaitGroup を利用して goroutine の完了を待つようにしたい
	wg := &sync.WaitGroup{}

	fmt.Println("== start ==")

	wg.Add(1)
	go task(ctx, wg)
	wg.Wait()

	fmt.Println("== finish ==")
}

func task(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	n := 5
	//  想定ではループ数は 3回で終了する
	for i := 0; i < n; i++ {
		select {
		case <-time.After(2 * time.Second):
			fmt.Printf("Done task (%v)!\n", i)
		case <-time.After(3 * time.Second):
			fmt.Printf("Done task (%v)!\n", i)
		case <-ctx.Done():
			fmt.Printf("cannel the context (%v).\n", i)
			return ctx.Err()
		}
	}
	return nil
}

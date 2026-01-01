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
	// task の戻り値を受け取るチャンネルを準備
	errCh := make(chan error, 1)

	fmt.Println("== start ==")

	wg.Add(1)
	// エラーを受け取れるようにする
	go func() {
		errCh <- task(ctx, wg)
	}()
	wg.Wait()

	if err := <-errCh; err != nil {
		fmt.Println("task error:", err)
	}

	fmt.Println("== finish ==")
}

func task(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	n := 5
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	// タイムアウトにより最後まで回り切らないことを確認する
	for i := 0; i < n; i++ {
		select {
		case <-ticker.C:
			fmt.Printf("Done task (%v)!\n", i)
		case <-ctx.Done():
			fmt.Printf("cancel the context (%v).\n", i)
			return ctx.Err()
		}
	}
	return nil
}

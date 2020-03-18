package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	wg := new(sync.WaitGroup)

	fmt.Println("== Start ==")

	wg.Add(1)
	go work(ctx, wg)
	wg.Wait()

	fmt.Println("== Finish ==")
}

func work(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	n := 5
	for i := 0; i < n; i++ {
		select {
		case <-time.After(2 * time.Second):
			fmt.Printf("Done work (%v)!\n", i)
		case <-ctx.Done():
			fmt.Printf("cannel the context (%v).\n", i)
			return ctx.Err()
		}
	}
	return nil
}

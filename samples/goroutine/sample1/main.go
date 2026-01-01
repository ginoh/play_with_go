package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		// goroutine の終了を待てるようにする
		defer wg.Done()
		fmt.Println("This is test")
	}()

	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())

	wg.Wait()
}

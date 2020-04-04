package main

import (
	"fmt"
	"runtime"
)

func main() {
	go fmt.Println("Yeah!")
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
}

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(2 * time.Second)
	end := time.Now()

	// end - start
	fmt.Printf("%v\n", end.Sub(start))
}

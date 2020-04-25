package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "key", "aaa")

	result := make(chan string)
	go SampleFunc(ctx, result)

	fmt.Printf("returned value: %s\n", <-result)

}

func SampleFunc(ctx context.Context, result chan<- string) error {
	v, ok := ctx.Value("key").(string)
	if !ok {
		fmt.Println("type assersion fail")
		// エラー処理
	}
	result <- (v + "bbb")
	return nil
}

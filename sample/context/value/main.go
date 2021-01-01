package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "key", "Hello")

	result := make(chan string)
	go SampleFunc(ctx, result)

	fmt.Printf("returned value: %s\n", <-result)

}

func SampleFunc(ctx context.Context, result chan<- string) error {
	v, ok := ctx.Value("key").(string)
	if !ok {
		fmt.Println("type assersion fail")

		// エラー処理などを書く
	}
	result <- (v + ", World!")
	return nil
}

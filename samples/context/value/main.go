package main

import (
	"context"
	"fmt"
)

type ctxKey string

const messageKey ctxKey = "message"

type Result struct {
	Value string
	Err   error
}

func main() {
	ctx := context.WithValue(context.Background(), messageKey, "Hello")

	// 結果とエラーを1本のチャネルで受け取る方式
	result := make(chan Result, 1)
	go SampleFunc(ctx, result)

	r := <-result
	if r.Err != nil {
		fmt.Println("error:", r.Err)
		return
	}
	fmt.Printf("returned value: %s\n", r.Value)

}

func SampleFunc(ctx context.Context, result chan<- Result) {
	v, ok := ctx.Value(messageKey).(string)
	if !ok {
		result <- Result{Err: fmt.Errorf("type assertion fail")}
		return
	}
	result <- Result{Value: v + ", World!"}
}

package main

import (
	"fmt"

	"./hoge"
)

func main() {
	fmt.Println(AppEcho())
	fmt.Println(hoge.FooEcho())
}

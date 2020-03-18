package main

import (
	"fmt"
)

func main() {
	m := make(map[int]string)

	m[1] = "hoge"
	m[10] = "foo"
	m[100] = "bar"

	fmt.Printf("%v\n", m)
}

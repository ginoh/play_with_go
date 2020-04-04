package main

import (
	"fmt"
)

var S = ""

func init() {
	S = S + "A"
}

func init() {
	S = S + "B"
}

func main() {
	fmt.Println(S)
}

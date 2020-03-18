package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}

	for i, v := range arr {
		fmt.Printf("%v: %v\n", i, v)
	}

	str := "Hello, World!"
	for i, r := range str {
		fmt.Printf("%v: %v, %s\n", i, r, string(r))
	}
	// indexの増分に注意
	str2 := "こんにちわ"
	for i, r := range str2 {
		fmt.Printf("%v: %v\n", i, r)
	}

}

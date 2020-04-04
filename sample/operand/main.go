package main

import "fmt"

func main() {
	a := 1
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", a<<1)
	fmt.Printf("%v\n", a<<2)

	b := 16
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", b>>1)
	fmt.Printf("%v\n", b>>2)
}

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

	m2 := map[int]string{1: "hoge", 10: "foo", 100: "bar"}
	fmt.Printf("%v\n", m2)

	v1 := m[1] // hoge
	fmt.Printf("%v\n", v1)
	v2, ok := m[2] // v2 == "", ok == false
	fmt.Printf("%#v, %#v\n", v2, ok)

	delete(m, 100)
	fmt.Printf("%v\n", m)
}

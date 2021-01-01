package main

import (
	"fmt"
)

func main() {
	a := [3]int{1, 2, 3}
	pow1(a)
	// 変化なし
	fmt.Printf("%v\n", a)

	b := []int{1, 2, 3}
	pow2(b)
	// 値の書き換え
	fmt.Printf("%v\n", b)

	c := make([]int, 3)
	c[0] = 5
	fmt.Printf("%v\n", c)

}

func pow1(a [3]int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}

// 参照渡し
func pow2(a []int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}

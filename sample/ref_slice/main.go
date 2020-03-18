package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	pow1(a)
	fmt.Printf("%v\n", a)

	s := []int{1, 2, 3}
	pow2(s)
	fmt.Printf("%v\n", s)

	b := make([]int, 3)
	b[0] = 5
	fmt.Printf("%v\n", b)

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

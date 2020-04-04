package main

import (
	"fmt"
)

func main() {
	//var s []int

	s := make([]int, 10)
	fmt.Println(s)
	s[0] = 10
	fmt.Println(s)

	s[1] = 9
	fmt.Println(s)

	fmt.Printf("%d\n", len(s))

	s2 := []int{1, 2, 3}
	s2 = append(s2, 4)
	fmt.Println(s2)

	s3 := []string{"Apple", "Banana", "Cherry"}

	for i, v := range s3 {
		fmt.Printf("[%d] => %s\n", i, v)
	}

	// appendによる追加分は処理対象にならない
	for i, v := range s3 {
		fmt.Printf("[%d] => %s\n", i, v)
		s3 = append(s3, "Melon")
	}
	fmt.Printf("%v\n", s3)

	fmt.Printf("%d\n", sum(1, 2, 3))

	s4 := []int{4, 5, 6}
	fmt.Printf("%d\n", sum(s4...))
}

func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

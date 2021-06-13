package main

import (
	"fmt"
)

func main() {
	s := make([]int, 5)
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

	// range を利用した for では
	// appendによる追加分は処理対象にならない
	for i, v := range s3 {
		fmt.Printf("[%d] => %s\n", i, v)
		s3 = append(s3, "Melon")
	}
	fmt.Printf("%v\n", s3)
	fmt.Printf("%d\n", sum(1, 2, 3))

	s4 := []int{4, 5, 6}
	// slice の値を展開して渡す
	fmt.Printf("%d\n", sum(s4...))
}

// 可変調引数
func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func push(s []int, v int) []int {
	return append(s, v)
}

func pop(s []int) (int, []int) {
	tail := s[len(s)-1]
	s = s[:len(s)-1]
	return tail, s
}

func unshift(s []int, v int) []int {
	t := make([]int, len(s), len(s)+1)
	copy(t, s)
	t = append(t[:1], t[0:]...)
	t[0] = v
	return t
}

func shift(s []int) (int, []int) {
	return s[0], s[1:]
}

func insert(s []int, v int, idx int) []int {
	t := make([]int, len(s), len(s)+1)
	copy(t, s)
	t = append(t[:idx+1], t[idx:]...)
	t[idx] = v
	return t
}

func delete(s []int, idx int) (int, []int) {
	del := s[idx]
	return del, append(s[:idx], s[idx+1:]...)
}

package main

import (
	"fmt"
)

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}

	// 5
	fmt.Println(arr1[4])

	arr1[4] = 6
	// 6
	fmt.Println(arr1[4])

	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{4, 5, 6}

	// 要素のコピー
	arr2 = arr3

	arr2[0] = 0

	fmt.Printf("%v\n", arr2)
	fmt.Printf("%v\n", arr3)
}

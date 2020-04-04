package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr[4])

	arr[4] = 6
	fmt.Println(arr[4])

	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{4, 5, 6}

	// 要素のコピー
	arr1 = arr2

	arr1[0] = 0

	fmt.Printf("%v\n", arr1)
	fmt.Printf("%v\n", arr2)
}

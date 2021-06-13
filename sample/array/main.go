package main

import "fmt"

func main() {
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{4, 5, 6}

	// 要素のコピー
	arr2 = arr3
	arr2[0] = 0
	fmt.Printf("%v\n", arr2)
	fmt.Printf("%v\n", arr3)

	// 2次元配列
	var arr [5][5]int

	num := 0
	for i, row := range arr {
		for j := range row {
			arr[i][j] = num
			num++
		}
	}
	// 配列の配列。ポインタの配列とかではない。
	fmt.Printf("%v\n", arr)
}

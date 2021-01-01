package main

import (
	"fmt"
)

func main() {
	// 2次元配列
	var arr [5][5]int

	num := 0
	for i, row := range arr {
		for j := range row {
			arr[i][j] = num
			num++
		}
	}
	fmt.Println(arr)
}

package main

import "fmt"

func main() {
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

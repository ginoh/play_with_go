package main

import (
	"fmt"
)

func main() {
	var str string
	// fmt.Scan は空白区切り
	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}

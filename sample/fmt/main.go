package main

import "fmt"

func main() {
	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}

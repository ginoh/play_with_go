package main

import "fmt"

func main() {
	var i int

	fmt.Printf("%T\n", &i)

	//var j *int
	// 初期化しなででリファレンスするとpanic
	//fmt.Println(*j)

	arr := &[3]int{1, 2, 3}

	fmt.Println((*arr)[0])
	// 上記はこれでも同じ結果を得られる
	fmt.Println(arr[0])

	var arr2 [3][3]int
	fmt.Printf("%T\n", arr2)

	arr3 := &[3]string{"hoge", "foo", "bar"}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}
}

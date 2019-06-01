package main

import (
	"fmt"
)

func main() {
	x := 1
	y := 5
	
	fmt.Println(plus(x, y))

	q, r := div(21, 5)

	fmt.Printf("quotient=%d remainder=%d\n", q, r)

	//戻り値を破棄するとき
	n, _ := div(10, 3)
	fmt.Println(n)

	//初期化された0が表示される
	fmt.Println(initValue())
}

// 関数名、引数の定義、戻り値の型定義
// 引数の型は同じものが連続する場合は最後まとめて定義できる
func plus (x, y int) int {
	return x + y
}
// 戻り値が複数ある時
func div (x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

//引数がなく、戻り値にint型のaが指定されている
func initValue() (a int) {
	return
}

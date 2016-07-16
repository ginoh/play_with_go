package main

import (
	"fmt"
)


//ラインコメント


/*
この書き方の時、途中にアスタリスクはいらないようだ
*/

//明示的な変数定義
var n int = 1

//暗黙の型定義はパッケージ変数ではできない？
//m := 10
var m = 10

func main() {
	//ローカル変数
	str := "test"

	//使われない変数があるとコンパイルエラーになる
	//noUse := "hoge"
		
	fmt.Printf("%d %s\n", n+m, str)

	//配列
	a := [3]int{1, 2, 3}
	fmt.Printf("%v\n", a)

	//要素数を初期化数から判定する
	b := [...]int{1, 2}
	fmt.Printf("%v\n", b)

	//interfaceというのがあるようだ
	//いろいろな型の値を代入できる
	//何に使えるかはまだよくわかってない
	var x interface{}
	x = 1
	x = "hoge"
	x = '茶'
	x = [...]int{100, 1000}
	fmt.Printf("%v\n", x)
}

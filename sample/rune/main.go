package main

import (
	"fmt"
)

func main() {
	// runeはUnicodeポイントを表す特殊な整数
	fmt.Println('a')
	fmt.Println('あ')
	fmt.Println('%')
	// バックスラッシュを利用するとUnicodeポイントそのもので表現できる
	fmt.Println('\000')
}

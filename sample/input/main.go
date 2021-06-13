package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	// fmt.Scan は空白区切り
	fmt.Println("== fmt.Scan ==")
	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)

	fmt.Println("== bufio.NewScanner ==")
	// 標準入力から入力を受け付ける scanner を生成
	scanner := bufio.NewScanner(os.Stdin)

	// 1行ずつ読み取り、EOFで終了
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		// non-EOF error
		fmt.Fprintln(os.Stderr, "エラー:", err)
	}
}

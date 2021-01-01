package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 標準入力から入力を受け付ける scanner を生成
	scanner := bufio.NewScanner(os.Stdin)

	// 1行ずつ読み取り、EOFで終了
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "エラー:", err)
	}
}

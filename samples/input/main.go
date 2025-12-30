package main

import (
	"bufio"
	"fmt"
	"io"
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

	fmt.Println("== bufio.NewReader ==")

	// 指定した長さのバイト列ごとに読み出す Reader 構造体を使う方法
	// 基本的には scanner 使っておくのが良さそう

	reader := bufio.NewReaderSize(os.Stdin, 4096)
	for {
		// バッファサイズを超えると isPrefix が true になり、バッファ分だけがまず返却される
		tmp, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			// non-EOF error
			fmt.Fprintln(os.Stderr, "エラー:", err)
			break
		}
		if !isPrefix {
			fmt.Println(string(tmp))
		}
	}
}

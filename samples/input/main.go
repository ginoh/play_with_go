package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var str string
	// 空白区切り
	fmt.Println("== fmt.Scan (Single token) ==")
	n, err := fmt.Scan(&str)
	if err != nil {
		fmt.Fprintln(os.Stderr, "scan error:", err)
		return
	}
	if n != 1 {
		fmt.Fprintln(os.Stderr, "scan: expected 1 token")
		return
	}
	fmt.Println(str)
	// 同じ標準入力を使うため、以降の例は入力の消費順に注意

	fmt.Println("== fmt.Scan (Multi token) ==")
	var str1, str2 string
	n, err = fmt.Scan(&str1, &str2)
	if err != nil {
		fmt.Fprintln(os.Stderr, "scan error:", err)
		return
	}
	if n != 2 {
		fmt.Fprintln(os.Stderr, "scan: expected 2 tokens")
		return
	}
	fmt.Println(str1, str2)

	fmt.Println("== fmt.Fscan ==")
	// io.Reader を指定して読み取れる
	fscanSample := "foo bar"
	fscanReader := strings.NewReader(fscanSample)
	var a, b string
	n, err = fmt.Fscan(fscanReader, &a, &b)
	if err != nil {
		fmt.Fprintln(os.Stderr, "fscan error:", err)
		return
	}
	if n != 2 {
		fmt.Fprintln(os.Stderr, "fscan: expected 2 tokens")
		return
	}
	fmt.Println(a, b)

	fmt.Println("== bufio.NewScanner ==")
	// 標準入力から入力を受け付ける scanner を生成
	scanner := bufio.NewScanner(os.Stdin)
	// 長い行に備えてバッファ上限を拡張
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	// 1行ずつ読み取り、EOFで終了
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		// non-EOF error
		fmt.Fprintln(os.Stderr, "エラー:", err)
	}

	fmt.Println("== bufio.Scanner (ScanWords) ==")
	sample := "Hello world 123"
	wordScanner := bufio.NewScanner(strings.NewReader(sample))
	wordScanner.Split(bufio.ScanWords)
	for wordScanner.Scan() {
		fmt.Println(wordScanner.Text())
	}
	if err := wordScanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "エラー:", err)
	}

	fmt.Println("== bufio.NewReader ==")

	// 指定した長さのバイト列ごとに読み出す Reader 構造体を使う方法
	// 基本的には scanner 使っておくのが良さそう
	// 1行読みの用途なら ReadString('\n') の方が扱いやすい
	reader := bufio.NewReaderSize(os.Stdin, 4096)
	var lineBuf bytes.Buffer
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
		lineBuf.Write(tmp)
		if isPrefix {
			continue
		}
		fmt.Println(lineBuf.String())
		lineBuf.Reset()
	}
}

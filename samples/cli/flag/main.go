package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		num    int
		output string
		dryrun bool
	)
	// 使用例を明示して help を分かりやすくする
	// -h/-helpでUsageが呼ばれたとき、フラグが存在しないなどの時に呼ばれる
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s -n 2 -o /tmp -dryrun\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.IntVar(&num, "n", 1, "client数")
	flag.StringVar(&output, "o", "", "出力先")
	// ロングオプションだけど、ダッシュ一つで表現される
	flag.BoolVar(&dryrun, "dryrun", false, "dryrun")

	flag.Parse()

	// 必須パラメータのチェック例
	if output == "" {
		fmt.Fprintln(os.Stderr, "出力先は必須です (-o)")
		os.Exit(1)
	}

	fmt.Println("client数=", num)
	fmt.Println("出力先=", output)
	fmt.Println("dryrun=", dryrun)

	// dryrun 時は実処理を行わない
	if dryrun {
		fmt.Println("dryrun のため処理をスキップします")
		return
	}
}

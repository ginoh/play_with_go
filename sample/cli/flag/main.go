package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		max    int
		output string
		dryrun bool
	)
	flag.IntVar(&max, "n", 1, "client数")
	flag.StringVar(&output, "o", "", "出力先")
	// ロングオプションだけど、ダッシュ一つ
	flag.BoolVar(&dryrun, "dryrun", false, "dryrun")

	flag.Parse()

	fmt.Println("client数=", max)
	fmt.Println("出力先=", output)
	fmt.Println("dryrun=", dryrun)
}

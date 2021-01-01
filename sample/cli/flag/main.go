package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		num    int
		output string
		dryrun bool
	)
	flag.IntVar(&num, "n", 1, "client数")
	flag.StringVar(&output, "o", "", "出力先")
	// ロングオプションだけど、ダッシュ一つで表現される
	flag.BoolVar(&dryrun, "dryrun", false, "dryrun")

	flag.Parse()

	fmt.Println("client数=", num)
	fmt.Println("出力先=", output)
	fmt.Println("dryrun=", dryrun)

	// go run main.go -n 2 -o /tmp -dryrun
}

package main

import (
	"fmt"

	"pkgtest2/app"
)

func main() {
	// module 内の別パッケージを呼び出す例
	str := app.AppEcho()
	fmt.Println(str)
}

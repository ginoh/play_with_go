package main

import (
	"fmt"

	"pkgtest3/app"
)

func main() {
	myApp := app.NewMyApp("MyApp")
	fmt.Println(myApp.Name())
}

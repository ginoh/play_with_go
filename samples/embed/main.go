package main

import (
	"embed"
	"fmt"
	"log"
)

//go:embed sources
var assets embed.FS

func main() {
	listDir("sources", "./dst")
}

func listDir(srcDir, dstDir string) {
	entries, err := assets.ReadDir("sources")
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range entries {
		fmt.Printf("src=%s\n", entry.Name())
	}
}

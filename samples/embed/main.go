package main

import (
	"embed"
	"fmt"
	"log"
	"path"
)

//go:embed sources
var assets embed.FS

func main() {
	listDirRecursive("sources")
}

func listDirRecursive(dir string) {
	entries, err := assets.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range entries {
		// embed FS 内の相対パスを組み立てる
		fullPath := path.Join(dir, entry.Name())
		if entry.IsDir() {
			listDirRecursive(fullPath)
			continue
		}
		fmt.Printf("src=%s\n", fullPath)
	}
}

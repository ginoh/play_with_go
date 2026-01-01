package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func main() {
	input := `{"id":1,"name":"Alice","tags":["a","b"],"active":true}`

	var m map[string]any
	if err := json.Unmarshal([]byte(input), &m); err != nil {
		log.Fatal(err)
	}

	// 型アサーションで取り出す
	if id, ok := m["id"].(float64); ok {
		fmt.Printf("id: %d\n", int(id))
	}
	if name, ok := m["name"].(string); ok {
		fmt.Printf("name: %s\n", name)
	}
	if active, ok := m["active"].(bool); ok {
		fmt.Printf("active: %v\n", active)
	}
	if tags, ok := m["tags"].([]any); ok {
		fmt.Printf("tags: %v\n", tags)
	}

	// Decoder を使った場合も同じ型になる
	dec := json.NewDecoder(strings.NewReader(input))
	var m2 map[string]any
	if err := dec.Decode(&m2); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("decoded: %#v\n", m2)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Value  string
	Before string
	Next   string
}

var (
	data   = make(map[string]*Node)
	oldest string
	latest string
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		params := strings.Split(text, " ")

		switch params[0] {
		case "add":
			add(params[1], params[2])
		case "get":
			get(params[1])
		case "remove":
			remove(params[1])
		case "evict":
			evict()
		case "exit":
			exit()
			return
		}
	}
}

func NewNode(value string) *Node {
	return &Node{
		Value:  value,
		Before: "",
		Next:   "",
	}
}

func add(key, value string) {
	if oldest == "" {
		data[key] = NewNode(value)
		oldest = key
		latest = key
		return
	}
	// 上書き
	if v, ok := data[key]; ok {
		data[key].Value = value

		if latest == key {
			return
		}

		if oldest == key {
			data[v.Next].Before = ""
			latest = key
			return
		}

		data[v.Before].Next = data[v.Next].Value
		data[v.Next].Before = data[v.Before].Value
		latest = key
		return
	}
	// 新規
	data[key] = NewNode(value)
	data[key].Before = latest
	data[latest].Next = key
	latest = key
}

func get(key string) {
	if v, ok := data[key]; ok {
		fmt.Println(data[key].Value)

		if latest == key {
			return
		}

		if oldest == key {
			data[v.Next].Before = ""
			oldest = v.Next
			data[latest].Next = key
			v.Next = ""
			latest = key
			return
		}

		data[v.Before].Next = data[v.Next].Value
		data[v.Next].Before = data[v.Before].Value
		latest = key
		return
	}
	fmt.Println("-1")
	return
}

func remove(key string) {
	// 削除
	if v, ok := data[key]; ok {
		if latest == key && oldest == key {
			latest = ""
			oldest = ""
			deleted := v.Value
			delete(data, key)
			fmt.Println(deleted)
			return
		}

		if latest == key {
			data[v.Before].Next = ""
			latest = v.Before
			deleted := v.Value
			delete(data, key)
			fmt.Println(deleted)
			return
		}

		if oldest == key {
			data[v.Next].Before = ""
			oldest = v.Next
			deleted := v.Value
			delete(data, key)
			fmt.Println(deleted)
			return
		}

		data[v.Before].Next = data[v.Next].Value
		data[v.Next].Before = data[v.Before].Value
		deleted := v.Value
		delete(data, key)
		fmt.Println(deleted)
		return
	}
	fmt.Println("-1")
}

func evict() {
	if oldest == "" {
		return
	}

	if v, ok := data[oldest]; ok {
		key := v.Next
		if key != "" {
			data[key].Before = ""
		}
		delete(data, oldest)
		oldest = key
	}
}

func exit() {
	fmt.Println("Good Bye!")
}

func getOldest() string {
	return oldest
}

func getLatest() string {
	return latest
}

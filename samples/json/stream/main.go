package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

func main() {
	// Decoder: JSONをストリームとして順次処理する例
	input := "{\"id\":1,\"name\":\"Alice\"}\n{\"id\":2,\"name\":\"Bob\"}"
	dec := json.NewDecoder(strings.NewReader(input))

	for {
		var u User
		if err := dec.Decode(&u); err != nil {
			if err == io.EOF {
				// 入力の終端なので正常終了
				break
			}
			log.Fatal(err)
		}
		fmt.Printf("decoded: %+v\n", u)
	}

	// Encoder: 構造体をストリームとして出力する例
	// io.Discard: https://pkg.go.dev/io#Discard
	enc := json.NewEncoder(io.Discard)
	if err := enc.Encode(User{ID: 3, Name: "Carol"}); err != nil {
		log.Fatal(err)
	}
}

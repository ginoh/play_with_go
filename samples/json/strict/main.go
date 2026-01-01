package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// 未知フィールドを許容しない例
	input := `{"id":1,"name":"Alice","unknown":"value"}`
	dec := json.NewDecoder(strings.NewReader(input))
	dec.DisallowUnknownFields()

	var u User
	if err := dec.Decode(&u); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", u)
}

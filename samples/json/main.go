package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

func main() {
	// struct -> JSON
	u := User{ID: 1, Name: "Alice"}
	b, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	// JSON -> struct
	var u2 User
	if err := json.Unmarshal(b, &u2); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", u2)
}

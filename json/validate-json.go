package main

import (
	"encoding/json"
	"fmt"
)

// let's use the json.Valid() method to check for malformed JSON in Go
func main() {
	good := `{"name": "John Doe"}`
	bad := `{name: "John Doe"}`

	fmt.Println(json.Valid([]byte(good)))
	fmt.Println(json.Valid([]byte(bad)))
}

/*
func Valid(data []byte) bool
*/

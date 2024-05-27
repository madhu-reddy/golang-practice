package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type Message struct {
		Name string
		Body string
		Time int64
	}

	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}

	fmt.Println(b)
	fmt.Println(string(b))
	fmt.Printf("Type of b: %T\n", b)
}

/*
func Marshal(v interface{}) ([]byte, error)
Output
[123 34 78 97 109 101 34 58 34 65 108 105 99 101 34 44 34 66 111 100 121 34 58 34 72 101 108 108 111 34 44 34 84 105 109 101 34 58 49 50 57 52 55 48 54 51 57 53 56 56 49 53 52 55 48 48 48 125]
{"Name":"Alice","Body":"Hello","Time":1294706395881547000}
Type of b: []uint8
*/

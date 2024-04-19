package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//JSON Marshal Int
	a, _ := json.Marshal(28192)
	fmt.Println(string(a)) // 20192

	a, _ = json.Marshal(true)
	fmt.Println(string(a)) // true

	// JSON Marshal Slice (For slices we will get a JSON array instead)
	a, _ = json.Marshal([]string{"foo", "bar", "baz"})
	fmt.Println(string(a)) // ["foo","bar","baz"]

	// JSON Marshal Map (Go maps will be marhsalled into JSON objects where keys of the map are keys of the JSON objects and values mapped to JSON object values)
	a, _ = json.Marshal(map[string]int{"foo": 1, "bar": 2, "baz": 3})
	fmt.Println(string(a)) // {"bar":2,"baz":3,"foo":1}44

}


/*

Here’s the definition of the json Marshal in the json encoding package

func Marshal(v interface{}) ([]byte, error)

We can effectively encode any type and we will get a byte slice as output and an error if there was any during the json marshalling.

*/




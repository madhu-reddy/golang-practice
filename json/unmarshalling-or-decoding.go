package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	input := `{
        "name": "John Doe",
        "age": 15,
        "hobbies": ["climbing", "cycling", "running"]
    }`

	var target map[string]any

	err := json.Unmarshal([]byte(input), &target)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}

	for k, v := range target {
		fmt.Printf("k: %s, v: %v\n", k, v)
	}

}

/*
The map[string]any data type in Go is a generic container that can hold values of any type, including complex nested structures.
In this case, the JSON keys will be unmarshalled into the string type, and their corresponding values will be unmarshalled into the any type of the map[string]any.

If an invalid JSON object is provided, an error will be returned by Unmarshal(). A common example of an invalid JSON object is one that has a trailing comma:
input := `
  {
    name: "John Doe",
    "age": 15,
    "hobbies": ["climbing", "cycling", "running"],
  }
`
*/

/*
Output
k: age, v: 15
k: hobbies, v: [climbing cycling running]
k: name, v: John Doe
*/

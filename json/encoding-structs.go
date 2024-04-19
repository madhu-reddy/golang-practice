package main

import (
	"encoding/json"
	"fmt"
)

// JSON Marshal Struct
type User struct {
	Name        string
	Age         int
	Active      bool
	lastLoginAt string
}

// JSON Marshal Struct with JSON Tags
type UserJsonTags struct {
	Name        string `json:"full_name"`
	Age         int    `json:"age,omitempty"`
	Active      bool   `json:"-"`
	lastLoginAt string
}

func main() {

	u, err := json.Marshal(User{Name: "Bob", Age: 10, Active: true, lastLoginAt: "today"})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(u)) // {"Name":"Bob","Age":10,"Active":true}

	user, err := json.Marshal(UserJsonTags{Name: "Bob", Age: 0, Active: true, lastLoginAt: "today"})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(user)) // {"full_name":"Bob"}

}

/* For "JSON Marshal Struct"
This is a more complete example and it allows you to see how a struct is being JSON marshalled in Go. All the exported struct fields.
Note that the field starting with lowercase lastLoginAt is unexported and wont be shown in the final JSON marshal output.
*/


/*
For "JSON Marshal Struct with JSON Tags"
You can specify a JSON tag json:"full_name" and specify a custom name for your JSON field
When specifying json:"<field name>,omitempty" as a json tag, the field will be disarcarded in the JSON output if the value has a zero value. This is mostly used to discard optional JSON fields in the output
When specifying json:"-" as a json tag, the field will be removed altogether from the JSON output. This is used when you want to keep the field in your Go struct but not in your JSON output
*/


/*
The Go standard library offers a package that has all you need to perform JSON encoding and decoding. The encoding/json package.
It allows you to do encoding of JSON data as defined in the RFC 7159.
When working with data we usually want to encode some Go struct into a json string.
However the Go standard library encoding/json package allows us to encode any data type into a JSON string.

Here’s the definition of the json Marshal in the json encoding package

func Marshal(v interface{}) ([]byte, error)

We can effectively encode any type and we will get a byte slice as output and an error if there was any during the json marshalling.

*/

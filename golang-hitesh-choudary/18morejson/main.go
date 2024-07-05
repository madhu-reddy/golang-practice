package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []courses{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 399, "LearnCodeOnline.in", "hit123", nil}}

	// package this data as JSON data
	//finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n", finalJson)
	fmt.Println(string(finalJson))

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev", "js"]
		}
	`)

	var lcoCourse courses
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		//fmt.Println(lcoCourse)
		fmt.Printf("%#v\n", lcoCourse) // main.courses{Name:"ReactJS Bootcamp", Price:299, Platform:"LearnCodeOnline.in", Password:"", Tags:[]string{"web-dev", "js"}}
	} else {
		fmt.Println("JSON was not valid")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is %T\n", k, v, v)
	}
}

/*
Output
Welcome to JSON video
JSON was valid
main.courses{Name:"ReactJS Bootcamp", Price:299, Platform:"LearnCodeOnline.in", Password:"", Tags:[]string{"web-dev", "js"}}
map[string]interface {}{"Price":299, "coursename":"ReactJS Bootcamp", "tags":[]interface {}{"web-dev", "js"}, "website":"LearnCodeOnline.in"}
Key is coursename and value is ReactJS Bootcamp and Type is string
Key is Price and value is 299 and Type is float64
Key is website and value is LearnCodeOnline.in and Type is string
Key is tags and value is [web-dev js] and Type is []interface {}
*/

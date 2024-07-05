package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	PerformPostFormRequest()
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// fake form data
	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "choudary")
	data.Add("email", "hitesh@go.dev")
	data.Add("email", "hitesh@go.in")

	fmt.Println(data) // map[email:[hitesh@go.dev hitesh@go.in] firstname:[hitesh] lastname:[choudary]]

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content)) // {"email":["hitesh@go.dev","hitesh@go.in"],"firstname":"hitesh","lastname":"choudary"}

}

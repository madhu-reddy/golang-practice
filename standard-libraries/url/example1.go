package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.example.com:8080/path?param1=value1&param2=value2"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

}


/*
Output

Welcome to handling URLs in golang
https://www.example.com:8080/path?param1=value1&param2=value2
https
www.example.com:8080
/path
8080
param1=value1&param2=value2
*/




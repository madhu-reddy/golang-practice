package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Welcome to handling URLs in golang")

	partsofurl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=hitesh",
	}

	anotherurl := partsofurl.String()
	fmt.Println(anotherurl)

}

/*

Welcome to handling URLs in golang
https://lco.dev/tutcss?user=hitesh

*/



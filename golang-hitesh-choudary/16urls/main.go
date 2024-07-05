package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=hgb5678"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery) // coursename=reactjs&paymentid=hgb5678

	qparams := result.Query() // This parses RawQuery
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"]) // [reactjs]

	for key, value := range qparams {
		fmt.Println(key, value)
	}

}

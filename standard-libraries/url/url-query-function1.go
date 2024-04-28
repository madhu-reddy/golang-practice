package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	result, err := url.Parse(myurl) //https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb

	if err != nil {
		panic(err)
	}
	fmt.Println(result.RawQuery) //coursename=reactjs&paymentid=ghbj456ghb

	qparams := result.Query()
	fmt.Println(qparams["coursename"]) //[reactjs]

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
		/*
		Param is:  [reactjs]
		Param is:  [ghbj456ghb]
		*/
	}

}

package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	websitelist := []string{
		"https://google.com",
		"https://facebook.com",
		"https://github.com",
		"https://facebook.com",
		"https://go.dev",
		"https://scientiamobile.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
}

func getStatusCode(endpoint string) {

	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	}

	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
}

/*
If there are three URLs in the list, wg.Add(1) will be called three times. This means the WaitGroup counter will be incremented to 3.
*/

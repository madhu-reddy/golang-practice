package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup

var mut sync.Mutex

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
		wg.Add(1)
		go getStatusCodeTest(web)

	}

	wg.Wait()
	fmt.Println(signals)
}

func getStatusCodeTest(endpoint string) {

	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func madhura() {
	defer wg.Done()
	fmt.Println("Testing Go Routines with wait groups, instead of time.sleep")
	// wg.Done()

}

func jaggura() {
	fmt.Println("Testing 2 Go Routines with wait groups, instead of time.sleep")
	wg.Done()

}

func main() {

	fmt.Println("This is Main")

	wg.Add(2)
	go jaggura()
	go madhura()
	wg.Wait()

}

/*
Global Declaration:
var wg sync.WaitGroup allows wg to be accessed by both the main and madhu functions.

Using defer wg.Done():
This ensures that wg.Done() is called when the madhu function completes, even if there’s an error or panic.

wg.Add(2):
Adds to the counter, indicating that we are waiting for one goroutine.

wg.Wait():
Blocks until all added goroutines have called wg.Done().
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func madhura() {
	fmt.Println("Testing Go Routines with wait groups, instead of time.sleep")
	wg.Done()

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

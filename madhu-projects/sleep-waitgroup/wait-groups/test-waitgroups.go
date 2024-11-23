package main

import (
	"fmt"
	"sync"
)

// var wg sync.WaitGroup

func madhur(wg *sync.WaitGroup) {
	fmt.Println("Testing Go Routines with wait groups, instead of time.sleep")
	wg.Done()

}

func jaggu(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Testing 2 Go Routines with wait groups, instead of time.sleep")
	// wg.Done()

}

func main() {
	var wg sync.WaitGroup
	fmt.Println("This is Main")

	wg.Add(2)
	go jaggu(&wg)
	go madhur(&wg)
	wg.Wait()

}

/*
Declare wg Inside main:
var wg sync.WaitGroup is declared as a local variable in main. This is valid and local to main, but it will be shared with goroutines if passed by reference.

Pass wg by Reference:
In the madhu function, wg is passed as a pointer (*sync.WaitGroup). This allows the madhu function to modify the same sync.WaitGroup instance declared in main.

Use defer wg.Done():
This ensures wg.Done() is called when the goroutine completes, even if there are errors or panics.
*/

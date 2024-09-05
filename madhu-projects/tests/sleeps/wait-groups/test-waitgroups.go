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
	fmt.Println("Testing 2 Go Routines with wait groups, instead of time.sleep")
	wg.Done()

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
2024-09-05-10-29-50.png
*/

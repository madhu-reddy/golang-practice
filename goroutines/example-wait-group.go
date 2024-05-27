package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{} // empty struct
	// add, done, wait

	wg.Add(2)
	go func() {
		fmt.Println("Hello")
		wg.Done()
	}()

	go func() {
		fmt.Println("World")
		wg.Done()
	}()
	fmt.Println("test waitgroups")
	wg.Wait()
	fmt.Println("Fnished")
}

/*
type WaitGroup struct {
	// contains filtered or unexported fields
}
*/

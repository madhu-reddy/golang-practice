package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Main started")
	var wg sync.WaitGroup

	wg.Add(1)
	go madhu(&wg)
	wg.Wait()

}

func madhu(wg *sync.WaitGroup) {
	fmt.Println("This is Madhu function")
	wg.Done()
}

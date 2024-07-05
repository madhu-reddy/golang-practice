package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang- LearnCodeOnline.in")

	myCh := make(chan int)
	wg := &sync.WaitGroup{}

	//myCh <- 5
	//fmt.Println(<-myCh) ---> deadlock error

	//fmt.Println(<-myCh) ---> deadlock error
	//myCh <- 5

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-myCh)
		fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		myCh <- 5
		myCh <- 6
		wg.Done()
	}(myCh, wg)

	wg.Wait()

}

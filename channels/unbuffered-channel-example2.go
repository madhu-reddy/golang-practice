package main

import (
	"fmt"
	"time"
)

func access(ch chan int) {
	time.Sleep(time.Second)
	fmt.Println("start accessing channel\n")

	for i := range ch { // This range iterates over each element as it’s received from queue.
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int)
	defer close(ch)

	go access(ch)

	for i := 0; i < 9; i++ {
		ch <- i
		fmt.Println("Filled")
	}

	time.Sleep(3 * time.Second)
}

/*
Output
start accessing channel

0
Filled
Filled
1
2
Filled
3
Filled
4
Filled
5
Filled
6
Filled
7
Filled
8
Filled
*/

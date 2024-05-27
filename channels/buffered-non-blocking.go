package main

import "fmt"

func main() {
	msgch := make(chan int, 5)
	msgch <- 10 // non blocking, unlike buffered channels
	msg := <-msgch
	fmt.Println(msg) // 10
}

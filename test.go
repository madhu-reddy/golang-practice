package main

import (
	"fmt"
)

func main() {
	ch := make(chan int) // Unbuffered channel

	func() {
		fmt.Println("Sending")
		ch <- 42 // Send value (blocks until receiver is ready)
		fmt.Println("Sent")
	}()

	fmt.Println(<-ch) // Receive value (blocks until sender sends value)
	fmt.Println("Received")
}

package main

import (
	"fmt"
	"time"
)

func sender(ch chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Sending:", i)
		ch <- i // Send value into the channel
	}
	close(ch) // Close the channel when done
}

func receiver(ch <-chan int) {
	for {
		val, ok := <-ch // Receive value from the channel
		if !ok {
			fmt.Println("Channel closed")
			return
		}
		fmt.Println("Received:", val)
	}
}

func main() {
	ch := make(chan int) // Create an unbuffered channel

	go sender(ch)   // Start sender goroutine
	go receiver(ch) // Start receiver goroutine

	// Wait for goroutines to finish
	time.Sleep(6 * time.Second)
}


/*
Channels are the pipes that connect concurrent goroutines.
You can send values into channels from one goroutine and receive those values into another goroutine.
*/

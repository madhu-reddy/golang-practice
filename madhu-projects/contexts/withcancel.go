package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a context with a cancel function
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when main returns
	// Create a channel to receive the result of the long-running operation
	ch := make(chan string)

	// Start the long-running operation as a goroutine
	go func() {
		time.Sleep(5 * time.Second) // simulate a long-running operation
		ch <- "result"
	}()

	// Wait for the long-running operation to finish or be canceled
	select {
	case res := <-ch:
		fmt.Println(res)
	case <-ctx.Done():
		fmt.Println("main canceled")
	}
}

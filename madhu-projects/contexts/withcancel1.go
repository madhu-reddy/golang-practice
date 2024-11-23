package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a context with a timeout of 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a channel to receive task completion status
	doneChan := make(chan bool)

	// Start performing the task
	go performTask(ctx, doneChan)

	// Listen for task completion or timeout
	select {
	case <-ctx.Done():
		fmt.Println("Task timed out")
	case success := <-doneChan:
		if success {
			fmt.Println("Task completed successfully")
		} else {
			fmt.Println("Task failed")
		}
	}
}

func performTask(ctx context.Context, doneChan chan<- bool) {
	select {
	case <-time.After(5 * time.Second):
		// Simulate a task taking 5 seconds
		doneChan <- true // Task completed successfully
	case <-ctx.Done():
		// Context was canceled, signal task failure
		doneChan <- false
	}
}

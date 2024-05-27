package main

import (
	"sync"
	"time"
)

func main() {
	c := make(chan string)
	defer close(c)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		c <- `foo` // This first goroutine is blocked after sending the message "foo" untill some receiver is ready to consume the channel data.
	}()

	go func() {
		defer wg.Done()

		time.Sleep(time.Second * 1) // waiting 1 second to ensure the other goroutine writes the data to the channel, before we read from the channel in this goroutine.
		//println(`Message: ` + <-c) -----> Result in an error if commented out "all goroutines are asleep - deadlock!"
		println(`Message: ` + <-c) // Message: foo
	}()

	wg.Wait()
}

/*
Unbuffered Channels:
An unbuffered channel is a channel that needs a receiver as soon as a message is emitted to the channel.
Sending or receiving on an unbuffered channel blocks the goroutine until the other side is ready.
This is useful for ensuring that two goroutines synchronize on some event.
ch := make(chan int)  // Unbuffered channl
*/

package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 4; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}
func main() {

	// creates capacity of 2
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(10 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		//time.Sleep(2 * time.Second)

	}
}

/*
Buffered Channels: These have a buffer of a fixed size and will only block when the buffer is full (sending) or empty (receiving).
Unlike the Unbuffered Channel, Buffered Channel has a capacity to store messages inside it. Buffered Channel could be filled up to its defined capacity, not only one message.
*/

/*
Output
successfully wrote 0 to ch
successfully wrote 1 to ch
read value 0 from ch
read value 1 from ch
read value 2 from ch
successfully wrote 2 to ch
successfully wrote 3 to ch
read value 3 from ch
*/

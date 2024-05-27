package main

import "fmt"

func main() {
	msgch := make(chan int)
	msgch <- 10        // blocking and the execution never goes to the next line
	fmt.Println(msgch) // fatal error: all goroutines are asleep - deadlock!
}

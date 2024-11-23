package main

import (
	"fmt"
	"time"
)

func madhu() {
	fmt.Println("this is another func")
}
func main() {
	fmt.Println("Hello")
	go madhu()
	time.Sleep(time.Second)
}

/*
func Sleep(d Duration)
Sleep pauses the current goroutine for at least the duration d. A negative or zero duration causes Sleep to return immediately.

In Go, constants like time.Second, time.Minute, etc., are defined in the time package.
If you check the source code of the time package in the Go standard library, you'll see that these constants are explicitly declared as constants of type time.Duration.

Here’s the relevant excerpt from the source code (time.go in the time package):

const (
    Nanosecond  Duration = 1
    Microsecond          = 1000 * Nanosecond
    Millisecond          = 1000 * Microsecond
    Second               = 1000 * Millisecond
    Minute               = 60 * Second
    Hour                 = 60 * Minute
)

Nanosecond is defined as 1 of type Duration (which is effectively 1 nanosecond).
Microsecond is 1000 * Nanosecond, Millisecond is 1000 * Microsecond, and so on.
These constants are explicitly typed as Duration

*/

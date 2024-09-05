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

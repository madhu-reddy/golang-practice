package main

import (
	"fmt"
	"time"
)

func main() {
	go greetergo("Hello")
	go greetergo("World")
	time.Sleep(1 * time.Millisecond)

}

func greetergo(s string) {
	for i := 0; i < 6; i++ {
		fmt.Println(s)
	}

}

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Main started")
	wg.Add(1)
	go madhu()
	wg.Wait()

}

func madhu() {
	fmt.Println("This is Madhu function")
	wg.Done()
}

/*
var m1 madhu:
Above is a declaration of a variable m1 of type madhu. In Go, when you declare a variable with var, it automatically gets initialized to the zero value of that type.
For structs, the zero value is an empty struct literal, i.e., madhu{}. So, when you write var m1 madhu, Go implicitly assigns m1 the value of madhu{}.

m1 := madhu
Above is incorrect because Go treats madhu as a type, not a value. In Go, the := operator expects an expression on the right-hand side, but a type (like madhu) is not an expression, hence the error: madhu (type) is not an expression
*/

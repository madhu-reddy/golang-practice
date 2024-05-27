// Understanding Empty Interfaces in Golang
package main

import (
	"fmt"
)

func observe(i interface{}) {

	// using the format specifier
	// %T to check type in interface
	fmt.Printf("The type passed is: %T\n", i)

	// using the format specifier %#v
	// to check value in interface
	fmt.Printf("The value passed is: %#v \n", i)
	fmt.Println("-------------------------------------")
}

func main() {
	var value float64 = 15
	value2 := "GeeksForGeeks"
	observe(value)
	observe(value2)
}

/*
The empty interface is extremely useful when we are declaring a function with unknown parameters and data types.

Output

The type passed is: float64
The value passed is: 15
-------------------------------------
The type passed is: string
The value passed is: "GeeksForGeeks"
-------------------------------------

*/

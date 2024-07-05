package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class of pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is", ptr) ---> Value of pointer is <nil> , which means pointing to nothing

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of actual pointer is", ptr)  // 0xc0000a4010
	fmt.Println("Value of actual pointer is", *ptr) // 23

	*ptr = *ptr * 2
	fmt.Println("New value is", myNumber) //New value is 46

}

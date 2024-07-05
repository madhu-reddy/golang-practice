package main

import "fmt"

func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	fmt.Println("vim-go")
	Anything(2.44)
	Anything("Angad")
	Anything(struct{}{}) // Initializing struct with null value, its like person{}, where person is a struct (type person struct {})
	// struct{}{} is a literal value of the empty struct type.

	myMap := make(map[string]interface{})
	myMap["name"] = "L04DB"
	myMap["age"] = 10
	fmt.Println(myMap)
}

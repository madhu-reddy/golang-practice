package main

import "fmt"

// remove a value from slice based on index
func main() {

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses) // [reactjs javascript swift python ruby]

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) // ... is variadic arguments, unpack slice elements
	fmt.Println(courses)                                    // [reactjs javascript python ruby]

}

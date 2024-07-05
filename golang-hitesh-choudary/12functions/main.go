package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult := proAdder(4, 5, 6, 7, 8, 9, 10)
	fmt.Println("proResult is: ", proResult)
}

func greeter() {
	fmt.Println("Namstey from golang")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}
	return total
}

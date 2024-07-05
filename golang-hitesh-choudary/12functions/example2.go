package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")

	proResult, _ := proAdder(4, 5, 6, 7, 8, 9, 10)
	fmt.Println("proResult is: ", proResult)
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, value := range values {
		total += value
	}
	return total, "proResult is"
}

package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch Out"
	} else {
		result = "Exactly 10 login count"
	}
	fmt.Println(result)
}

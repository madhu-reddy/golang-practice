package main

import "fmt"

func main() {
	fmt.Println("using loops")
	rougueValue := 2

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}

		if rougueValue == 9 {
			fmt.Println("Entered 9, exiting loop")
			break
		}

		fmt.Println("Value is: ", rougueValue)
		rougueValue++

	}
lco:
	fmt.Println("LCO, jumping at Learncodeonline.in")
}

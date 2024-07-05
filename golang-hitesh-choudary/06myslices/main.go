package main

import "fmt"

func main() {
	fmt.Println("Welcome to the video on slices")
	var fruitList = []string{} // empty slice
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	var vegList = []string{"Tomato", "cabbage", "spinach"}
	fmt.Println("VegList is", vegList) //VegList is [Tomato cabbage spinach]

	vegList = append(vegList, "cauliflower", "onion")
	fmt.Println("New VegList is", vegList) // [Tomato cabbage spinach cauliflower onion]

	vegList = append(vegList[1:3])
	fmt.Println("New VegList is", vegList) // [cabbage spinach]

}

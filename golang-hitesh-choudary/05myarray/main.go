package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")
	var fruitlist [4]string
	fruitlist[0] = "Apple"
	fruitlist[1] = "Tomato"
	fruitlist[2] = "Peach"

	fmt.Println("Fruit list is", fruitlist)      // Fruit list is [Apple Tomato Peach ]
	fmt.Println("Fruit list is", len(fruitlist)) // Fruit list is 4
}

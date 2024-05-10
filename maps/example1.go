package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println(m)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

	// Create a new slice with length 5 and capacity 10
	mySlice := make([]int, 5, 10)
	fmt.Println(mySlice)

	s := make([]int, 10, 15)
	fmt.Println(s)

	p := make(map[int]string)
	fmt.Println(p)

	// Create a new map
	myMap := make(map[string]int)
	fmt.Println(myMap)

	myIntPointer := new(int)
	fmt.Println(*myIntPointer)
}

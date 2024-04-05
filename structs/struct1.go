package main

import (
	"fmt"
	"strconv"
)

type madhu struct {
	Name string
	age  int
}

func (m madhu) test() string {
	return m.Name + " age is " + strconv.Itoa(m.age)
}

func main() {
	tes := madhu{Name: "maddy", age: 20}
	fmt.Println(tes.test())
}


/*
output

maddy age is 20
*/

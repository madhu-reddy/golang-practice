package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func (p *person) agefunc() {
	p.age = 40
}
func main() {
	fmt.Println("Hello World!")
	p := person{name: "madhu", age: 37}
	fmt.Printf("Person name is '%v' and age is '%v'\n", p.name, p.age)
	p.agefunc()
	fmt.Printf("Person name is '%v' and age is '%v'\n", p.name, p.age)
}


/*
Output

Hello World!
Person name is 'madhu' and age is '37'
Person name is 'madhu' and age is '40'

*/



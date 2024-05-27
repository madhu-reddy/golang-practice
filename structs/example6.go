package main

import (
	"fmt"
)

type Car struct {
	Name    string
	Age     int
	ModelNo int
}

func (c Car) Print() {
	fmt.Println(c)
}

func (c Car) Drive() {
	fmt.Println("driving...")
}

func (c Car) GetName() string {
	return c.Name

}

func main() {
	c := Car{
		Name:    "Chevy",
		Age:     1,
		ModelNo: 2,
	}
	c.Print()
	c.Drive()
	fmt.Println(c.GetName())
}

/*
Output
{Chevy 1 2}
driving...
Chevy
*/

package main

import "fmt"

type Car struct {
	Name    string
	Age     int
	ModelNo int
}

func (c Car) Print() {
	fmt.Println(c)
	fmt.Println(c.Age)
}

func (c Car) GetName() string {
	return c.Name

}

func main() {
	c := Car{
		Name:    "chevy",
		Age:     1,
		ModelNo: 2,
	}

	//fmt.Println(c)
	c.Print()
	fmt.Println(c.GetName())
}

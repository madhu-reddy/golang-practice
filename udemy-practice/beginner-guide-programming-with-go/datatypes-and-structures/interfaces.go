package main

import (
	"fmt"
)

type Car interface {
	Drive()
	Stop()
}

func NewModel(arg string) Car {
	return &Lambo{arg}
}

type Lambo struct {
	LamboModel string
}

type Chevy struct {
	ChevyModel string
}

// Lambo struct implements/follows the "Car" Interface, as it implements both the Interface methods.
func (l *Lambo) Drive() {
	fmt.Println("Lambo on the move")
	fmt.Println(l.LamboModel)
}

func (l *Lambo) Stop() {
	fmt.Println("Stopping Lambo")
}

// Chevy struct  doesn't implements/follows the "Car" Interface
func (c *Chevy) Drive() {
	fmt.Println("Chevy on the move")
	fmt.Println(c.ChevyModel)
}

func main() {
	l := Lambo{"Gallardo"}
	l1 := NewModel("Gallardo")

	c := Chevy{"C369"}

	l.Drive()
	l1.Drive()

	c.Drive()

}

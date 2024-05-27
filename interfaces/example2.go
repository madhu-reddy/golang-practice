package main

import (
	"fmt"
)

type Car interface {
	Drive()
	Stop()
}

type Lambo struct {
	LamboModel string
}

type Chevy struct {
	ChevyModel string
}

func NewModel(arg string) Car {
	return &Lambo{arg}
}

func (l *Lambo) Drive() {
	fmt.Println("Lambo on the move")
	fmt.Println(l.LamboModel)
}

/*func (l *Lambo) Stop() {
	fmt.Println("Stopping Lambo")

}*/ // cannot use &Lambo{…} (value of type *Lambo) as Car value in return statement: *Lambo does not implement Car (missing method Stop)

func (l *Lambo) Stop() {
	fmt.Println("Stopping Lambo")

}

func (c Chevy) Drive() {
	fmt.Println("Chevy on the move")
	fmt.Println(c.ChevyModel)
}

func main() {
	c := Chevy{"C369"}
	c.Drive()
	l := NewModel("Gallardo")
	l.Drive()

}

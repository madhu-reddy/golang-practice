package main

import "fmt"

// interface
type school interface {
	class() string
}

// struct to implement interface
type madhu struct {
	name string
	age  int
	cls  string
}

type ramu struct {
	name string
	age  int
	cls  string
}

type manoj struct {
	name string
	age  int
	cls  string
}

// use struct to implement class() of interface
func (s madhu) class() string {
	return s.name + " belongs to class " + s.cls
}

// use struct to implement class() of interface
func (s ramu) class() string {
	return s.name + " belongs to class " + s.cls
}

func (s manoj) class() string {
	return s.name + " belongs to class " + s.cls
}

// access method of the interface
func intermethod(i school) {
	fmt.Println("Student info: ", i.class())
}

// main function
func main() {

	// assigns value to struct members
	user1 := madhu{name: "madhusudan", cls: "eighth"}
	user2 := ramu{name: "rammohan", cls: "tenth"}
	user3 := manoj{name: "manoj chowdary", cls: "fifth"}

	intermethod(user1)
	intermethod(user2)
	intermethod(user3)

}



/*
Output

Student info:  madhusudan belongs to class eighth
Student info:  rammohan belongs to class tenth
Student info:  manoj chowdary belongs to class fifth

*/


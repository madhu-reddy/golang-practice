package main

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

func main() {
	emp8 := &Employee{
		firstName: "Sam",
		lastName:  "Anderson",
		age:       55,
		salary:    6000,
	}

	fmt.Println("First Name:", (*emp8).firstName)
	fmt.Println("Age:", (*emp8).age)

	fmt.Println("First Name:", emp8.firstName)
	fmt.Println("Age:", emp8.age)

}

/*
The Go language gives us the option to use emp8.firstName instead of the explicit dereference (*emp8).firstName to access the firstName field as well as age field.
*/

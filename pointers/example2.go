package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name, age: 50}
	return &p
}

func main() {
	fmt.Println(newPerson("Jon").name)
	fmt.Println(newPerson("Jon").age)
	fmt.Println(newPerson("Jon"))
	fmt.Println(*newPerson("Jon"))

}

/*
Output

Jon
50
&{Jon 50}
{Jon 50}

*/

/*
fmt.Println(newPerson("Jon").name) vs fmt.Println(*newPerson("Jon").name)

fmt.Println(newPerson("Jon").name) accesses the name field directly from the person struct pointer without dereferencing it, while fmt.Println(*newPerson("Jon").name) attempts to dereference the pointer, which is invalid for non-pointer types like strings.
*newPerson("Jon").name is invalid, it's trying to dereference a string, which is not a pointer type.
 */

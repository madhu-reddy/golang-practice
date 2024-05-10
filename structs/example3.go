package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))
	fmt.Println(*newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}

/*
Output
{Bob 20}
{Alice 30}
{Fred 0}
&{Ann 40}
&{Jon 42}
{Jon 42}
Sean
50
51
{Rex true}
*/

/*
1) You can name the fields when initializing a struct.
2) Omitted fields will be zero-valued.
3) An & prefix yields a pointer to the struct.
4) Access struct fields with a dot.
5) Structs are mutable.
6) You can also use dots with struct pointers - the pointers are automatically dereferenced.
*/

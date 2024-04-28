package main

import "fmt"

func my_age(age int) {
	age = 20
}

func my_age_pointer(myageptr *int) {
	*myageptr = 40
}

func main() {
	age := 30
	fmt.Println("initial:", age)

	my_age(age)
	// "age" variable after my_age() function call
	fmt.Println("after my_age func call:", age)

	my_age_pointer(&age)
	// "age" variable after my_age_pointer() function call.
	fmt.Println("after my_age_pointer func call:", age)
	fmt.Println("after my_age_pointer func call:", *&age)

	// "age" variable memory address
	fmt.Println("address of age variable:", &age)

}

/*

initial: 30
after my_age func call: 30
after my_age_pointer func call: 40
after my_age_pointer func call: 40
address of age variable: 0xc0000120b0

*/



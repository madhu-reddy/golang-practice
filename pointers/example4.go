package main

import (
	"fmt"
)

func swap(m1 *int, m2 *int) {
	var temp int
	temp = *m2
	*m2 = *m1
	*m1 = temp
}

func main() {
	var a int = 10
	var b int = 20
	fmt.Printf("a = %d, b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)
}

/*
Output
a = 10, b = 20
a = 20, b = 10
*/

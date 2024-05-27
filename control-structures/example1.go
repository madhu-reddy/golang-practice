package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	f := true
	flag := &f

	if flag == nil {
		fmt.Println("Value is nill")
	} else if *flag {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	arr := []string{"angad", "sharma", "middlename"}
	for i, value := range arr {
		fmt.Println(i)
		fmt.Println(value)
	}

	mymap := make(map[string]interface{})
	mymap["name"] = "angad"
	mymap["age"] = 20
	for k, v := range mymap {
		fmt.Printf("Key: %s and Value: %v", k, v)
		fmt.Println()
	}

}

/*
Output
Hello World
True
0
angad
1
sharma
2
middlename
Key: name and Value: angad
Key: age and Value: 20
*/

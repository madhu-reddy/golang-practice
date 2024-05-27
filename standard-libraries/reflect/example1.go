package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 10
	name := "Go Lang"
	type Book struct {
		name   string
		author string
	}
	sampleBook := Book{"Reflection in Go", "John"}
	fmt.Println(reflect.TypeOf(x))                    // int
	fmt.Println(reflect.ValueOf(name))                // Go Lang
	fmt.Println(reflect.ValueOf(sampleBook))          // {Reflection in Go John}
	fmt.Println(reflect.ValueOf(sampleBook).Field(1)) // John
	fmt.Println(reflect.TypeOf(sampleBook))           // main.Book
}

/*
The ValueOf function returns a reflection Value instance based on the provided variable.
Similar to the reflection Type, reflection Value also holds more information about the variable’s value.
*/

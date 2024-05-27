package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "I'm a string!"
	sValue := reflect.ValueOf(s)
	fmt.Println(sValue.Len())       // 13
	fmt.Println(sValue.Interface()) // I'm a string!

	stringType := reflect.TypeOf(s) // or sValue.Type()
	fmt.Println(stringType.Name())  // "string"

	type myStruct struct {
		foo int
		bar bool
	}

	reflect.ValueOf(myStruct{123, false}).Field(0) // reflect.ValueOf(123)

}

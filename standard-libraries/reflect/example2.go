package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Book struct {
		Name   string
		Author string
		Year   int
	}
	sampleBook := Book{"Reflection in Go", "John", 2021}
	val := reflect.ValueOf(sampleBook)
	fmt.Println(val)

	for i := 0; i < val.NumField(); i++ {
		fieldName := val.Type().Field(i).Name
		fieldValue := val.Field(i).Interface()
		fmt.Println(fieldName, " -> ", fieldValue)
	}

}

/*
output
Name  ->  Reflection in Go
Author  ->  John
Year  ->  2021
*/

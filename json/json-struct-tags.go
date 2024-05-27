package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	type Dog struct {
		Breed         string
		Name          string
		FavoriteTreat string
		Age           int
	}

	var dog = Dog{
		Breed:         "Golden Retriever",
		Age:           8,
		Name:          "Paws",
		FavoriteTreat: "Kibble",
	}

	/* At the moment, the dog variable will be marshalled into the following JSON object where the properties correspond exactly to the struct field names:
	{"Breed":"Golden Retriever","Name":"Paws","FavoriteTreat":"Kibble","Age":8} */

	m, err := json.Marshal(dog)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(m)
	fmt.Println(string(m)) // {"Breed":"Golden Retriever","Name":"Paws","FavoriteTreat":"Kibble","Age":8}

	// You can customize this output by using json struct tags on the Dog type:
	type Dogtag struct {
		Breed         string `json:"breed"`
		Name          string `json:"name"`
		FavoriteTreat string `json:"favorite_treat"`
		Age           int    `json:"age"`
	}

	var dogtag = Dogtag{
		Breed:         "Golden Retriever",
		Age:           8,
		Name:          "Paws",
		FavoriteTreat: "Kibble",
	}

	mtag, err := json.Marshal(dogtag)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(mtag)) // {"breed":"Golden Retriever","name":"Paws","favorite_treat":"Kibble","age":8}
}

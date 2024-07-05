package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:", languages) // List of all languages: map[JS:Javascript PY:Python RB:Ruby]

	for key, value := range languages {
		fmt.Printf("For key is %v, value is %v\n", key, value)
	}

}

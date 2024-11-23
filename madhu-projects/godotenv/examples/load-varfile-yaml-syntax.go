package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	var myEnv map[string]string
	myEnv, err := godotenv.Read(".madhuyamlenv")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("myfoo: ", myEnv["FOO"])
	fmt.Println("mybar: ", myEnv["BAR"])

}

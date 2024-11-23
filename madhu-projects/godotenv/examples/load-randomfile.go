package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".madhuenv")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	somevar := os.Getenv("SOME_VAR")
	somefoo := os.Getenv("FOO")
	somebar := os.Getenv("BAR")

	// now do something with s3 or whatever

	fmt.Println(somevar)
	fmt.Println(somefoo)
	fmt.Println(somebar)
}

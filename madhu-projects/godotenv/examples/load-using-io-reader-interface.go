package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getLocalFile(filePath string) (*os.File, error) {
	// Open the local .env file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func main() {
	// Get the reader for the local .env file
	file, err := getLocalFile(".env") // Replace with your .env file path
	if err != nil {
		log.Fatalf("Failed to open local file: %v", err)
	}
	defer file.Close()

	// Parse the environment variables from the file
	myEnv, err := godotenv.Parse(file)
	if err != nil {
		log.Fatalf("Error parsing environment variables: %v", err)
	}

	// Print out the parsed environment variables
	for key, value := range myEnv {
		fmt.Printf("%s=%s\n", key, value)
	}
}

/*
os.File implements the Read() method of the io.Reader interface, which means an *os.File object is considered to be of type io.Reader
*/

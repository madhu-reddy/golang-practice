package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getLocalFileContent(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func main() {
	// Get the local .env file content
	content, err := getLocalFileContent(".env")
	if err != nil {
		log.Fatalf("Failed to read local file: %v", err)
	}

	// Parse the environment variables from the string content
	myEnv, err := godotenv.Unmarshal(content)
	if err != nil {
		log.Fatalf("Error parsing environment variables: %v", err)
	}

	// Print the parsed environment variables
	for key, value := range myEnv {
		fmt.Printf("%s=%s\n", key, value)
	}
}

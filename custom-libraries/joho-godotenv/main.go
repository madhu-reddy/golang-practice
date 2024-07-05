package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	envFilePath := "config/.env"

	// Load the .env file
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get environment variables
	appEnv := os.Getenv("APP_ENV")
	port := os.Getenv("PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	// Print the environment variables
	fmt.Printf("App Environment: %s\n", appEnv)
	fmt.Printf("Port: %s\n", port)
	fmt.Printf("Database Host: %s\n", dbHost)
	fmt.Printf("Database User: %s\n", dbUser)
	fmt.Printf("Database Password: %s\n", dbPass)
}

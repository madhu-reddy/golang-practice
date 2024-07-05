package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/sys/execabs"
)

func main() {
	// Define the command and arguments
	cmd := execabs.Command("echo", "Hello, World!")

	// Set the command's output to standard output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Command execution failed: %v", err)
	}

	fmt.Println("Command executed successfully")
}

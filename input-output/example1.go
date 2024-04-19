package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	str := "My String Data"
	reader := strings.NewReader(str)

	// Create a buffer to store the read bytes
	bytes := make([]byte, 1024)

	// Read from the reader into the bytes buffer
	n, err := reader.Read(bytes)
	if err != nil && err != io.EOF {
		panic(err)
	}

	// Write the read bytes to stdout
	if _, err := os.Stdout.Write(bytes[:n]); err != nil {
		panic(err)
	}

	fmt.Println() // Print a newline for formatting
}

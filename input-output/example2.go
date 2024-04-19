package main

import (
	"fmt"
	"os"
)

func main() {
	// Read all bytes from a file
	data, err := os.ReadFile("myfile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Write all bytes to a file
	err = os.WriteFile("myfilewrite.txt", data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

/*
func ReadFile(name string) ([]byte, error)
ReadFile reads the named file and returns the contents. A successful call returns err == nil, not err == EOF. Because ReadFile reads the whole file, it does not treat an EOF from Read as an error to be reported.


func WriteFile(name string, data []byte, perm FileMode) error
WriteFile writes data to the named file, creating it if necessary. If the file does not exist, WriteFile creates it with permissions perm (before umask);

*/




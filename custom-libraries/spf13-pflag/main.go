package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

func main() {
	// Define flags
	var (
		stringFlag string
		intFlag    int
		boolFlag   bool
	)

	// Bind flags to variables
	pflag.StringVarP(&stringFlag, "name", "n", "default", "a string flag")
	pflag.IntVarP(&intFlag, "age", "a", 0, "an int flag")
	pflag.BoolVarP(&boolFlag, "verbose", "v", false, "a bool flag")

	// Parse the flags
	pflag.Parse()

	// Use the flags
	fmt.Printf("Name: %s\n", stringFlag)
	fmt.Printf("Age: %d\n", intFlag)
	fmt.Printf("Verbose: %v\n", boolFlag)
}

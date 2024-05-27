package main

import (
	"fmt"
	"runtime"
)

func main() {
	numCPU := runtime.NumCPU()
	fmt.Printf("Number of logical CPUs: %d\n", numCPU)

	// Set the number of threads to use for parallel execution
	previous := runtime.GOMAXPROCS(numCPU)
	fmt.Printf("No of OS threads for parallel execution: %d\n", previous)

	// Your parallel computation code would go here
}

/*
Number of logical CPUs: 8
No of OS threads for parallel execution: 8
*/

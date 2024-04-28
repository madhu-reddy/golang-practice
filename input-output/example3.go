package main

import (
	"fmt"
)

func main() {
	word := []byte("testingbytes")
	value := 0

	for _, c := range word {
		value += int(c)
	}
	fmt.Println(word)
	fmt.Printf("%s: %d\n", word, value)
}

/*
output
[116 101 115 116 105 110 103 98 121 116 101 115]
testingbytes: 1317
*/

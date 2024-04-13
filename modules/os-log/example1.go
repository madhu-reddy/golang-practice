package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("sloth.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("error opening sloth.txt: %v", err)
	}
	defer file.Close()

	// Make a byte slice that's big enough to store a few words of the message
	// we're reading
	bytesRead := make([]byte, 60)

	// Now read some data, passing in our byte slice
	n, err := file.Read(bytesRead)
	if err != nil {
		log.Fatalf("error reading from sloth.txt: %v", err)
	}

	// Take a look at the bytes we copied into the byte slice
	log.Printf("We read \"%s\" into bytesRead (%d bytes)",
		string(bytesRead), n)
}

/*
os.OpenFile returns a pointer to File struct (*FIle).
The File stuct implements the interface type "Reader", which has a method called "Read".

type Reader interface {
	Read(p []byte) (n int, err error)
}

func (f *File) Read(b []byte) (n int, err error)

*/


/*
cat  sloth.txt
I'm a sloth and hibiscus flowers are tasty!
*/


/*
Outut
2024/04/13 00:23:50 We read "I'm a sloth and hibiscus flowers are tasty!" into bytesRead (43 bytes)
*/

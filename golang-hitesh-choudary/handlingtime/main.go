package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006"))  // 05-25-2024
	fmt.Println(presentTime.Format(time.UnixDate)) // Sat May 25 06:56:10 EDT 2024

	createdDate := time.Date(2023, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)                             // 2023-08-10 23:23:00 +0000 UTC
	fmt.Println(createdDate.Format("01-02-2006 Monday")) // 08-10-2023 Thursday
	str, _ := presentTime.MarshalText()
	fmt.Println(string(str))

}

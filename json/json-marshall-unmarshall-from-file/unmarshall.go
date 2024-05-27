package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// Define a struct to match the structure of each JSON object
type Person struct {
	Criteria      []string `json:"criteria"`
	JsCheckFuncts []string `json:"js_check_functs"`
	Name          string   `json:"name"`
	Webgl         string   `json:"webgl"`
}

func main() {
	// Read JSON file
	jsonFile, err := os.Open("test.json")
	if err != nil {
		log.Fatalf("Failed to open JSON file: %v", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	// Unmarshal JSON data into a slice of Person structs
	var people []Person
	err = json.Unmarshal(byteValue, &people)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON data: %v", err)
	}
	fmt.Println(people)

}

/*
Output
[{[id^=apple_iphone id*=subhw2g brand_name=Apple model_name=iPhone 2G] [_window.screen.height==480 _window.devicePixelRatio==1 _navigator.geolocation==false _window.history.pushState==false] iPhone 2G } {[id^=apple_ipad id*=subhwpro713 brand_name=Apple model_name=iPad Pro 13 (2024)] [_window.screen.height==1376 _window.devicePixelRatio==2] iPad Pro 13 2024 }
*/

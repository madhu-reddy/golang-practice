package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// CustomTime embeds time.Time and implements json.Marshaler
type CustomTime struct {
	time.Time
}

// MarshalJSON implements the json.Marshaler interface
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.Format("2006-01-02 15:04:05"))
}

func main() {
	ct := CustomTime{Time: time.Now()}
	jsonData, err := json.Marshal(ct)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData)) // Outputs: "2024-05-21 14:35:02" (or current time)
}

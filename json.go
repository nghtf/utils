package utils

import (
	"encoding/json"
	"fmt"
)

// Convert any to JSON string
func ToJSON(v any) (string, error) {
	b, err := json.MarshalIndent(&v, "", " ")
	if err != nil {
		return "", err
	}
	fmt.Println(string(b))
	return (string(b)), nil
}

// Print as JSON
func PrintAsJSON(v any) {
	b, err := json.MarshalIndent(&v, "", " ")
	if err != nil {
		fmt.Println("PrintAsJSON() error:", err)
	}
	fmt.Println(string(b))
}

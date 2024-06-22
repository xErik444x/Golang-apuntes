package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	myJsonString := `{"usuario":"xErik444x"}`
	var myStoredVariable map[string]interface{}
	json.Unmarshal([]byte(myJsonString), &myStoredVariable)
	for key, value := range myStoredVariable {
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}
}

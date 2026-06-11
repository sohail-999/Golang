package main

import (
	"fmt"
)

func main() {
	myMap := map[string]int{"apple": 5, "banana": 10}

	value, ok := myMap["apple"]

	if ok {
		fmt.Println("Value found:", value) // Output: Value found: 5
	} else {
		fmt.Println("Key not found")
	}

	value, ok = myMap["cherry"]

	if ok {
		fmt.Println("Value found:", value)
	} else {
		fmt.Println("Key not found") // Output: Key not found
	}
}

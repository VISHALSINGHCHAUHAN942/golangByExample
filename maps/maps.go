package main

import (
	"fmt"
	"maps"
)

func main() {

	// Create a new empty map using make()
	// The map has string keys and int values.
	m := make(map[string]int)

	// Assign key-value pairs
	m["k1"] = 7
	m["k2"] = 8

	fmt.Println("Map:", m)

	// Access and print the value of key "k1"
	v1 := m["k1"]
	fmt.Println("v1 is", v1)

	// Access and print the value of key "k2"
	v2 := m["k2"]
	fmt.Println("v2 is", v2)

	// Accessing a non-existent key returns the zero value (0 for int)
	fmt.Println("k3 value is", m["k3"])

	// Print the total number of key-value pairs in the map
	fmt.Println("Length of map:", len(m))

	// Demonstrate deleting a key from the map
	fmt.Println("Before deleting:", m)
	delete(m, "k1")
	fmt.Println("After deleting:", m)

	// Check if a specific key is present in the map
	// The second return value (isKeyPresent) is a boolean
	_, isKeyPresent := m["k2"]
	fmt.Println("Is 'k2' present?", isKeyPresent)

	// Clear all key-value pairs from the map
	clear(m)
	fmt.Println("After clearing map:", m)

	// Initialize a map directly without using make()
	n := map[string]int{"l2": 10, "l4": 30}
	fmt.Println("Map n:", n)

	// Create another map with the same structure
	n5 := map[string]int{"l2": 12, "l4": 30}

	// Compare two maps using maps.Equal() (Go 1.21+ feature)
	if maps.Equal(n, n5) {
		fmt.Println("n == n5")
	} else {
		fmt.Println("n != n5")
	}
}

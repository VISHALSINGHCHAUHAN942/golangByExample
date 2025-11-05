package main

import (
	"fmt"
	"sort"
)

func main() {
	// Range over a Slice or Array
	fmt.Println("Range over slices/array")
	nums := []int{10, 20, 30, 40, 50}
	for index, value := range nums {
		fmt.Println("index", index, "Value", value)
	}

	// If you don’t need the index, you can ignore it with _
	for _, value := range nums {
		fmt.Println("value:", value)
	}

	// Range over a Map
	fmt.Println("Range over maps")
	m := map[string]int{"apple": 5, "banana": 7, "pomegrandte": 9}
	for key, value := range m {
		fmt.Println("Key of map is", key, "Value:", value)
	}

	// Process to get sorted keys of map
	// since maps are unordered, we manually sort keys.
	keySlice := make([]string, 0, len(m))
	for index := range m {
		keySlice = append(keySlice, index)
	}

	sort.Strings(keySlice)

	fmt.Println("Sorted map keys are:")
	for _, iterator := range keySlice {
		fmt.Println("Sorted key:", iterator, "Value:", m[iterator])
	}

	// Range over string
	var s string = "Golang"

	// Iterate over string using range
	// `char` is a rune (Unicode code point).
	// Converting to string prints the actual character.
	for index, char := range s {
		fmt.Println("Index:", index, "Character:", string(char))
	}

	
}

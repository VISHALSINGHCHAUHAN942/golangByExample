package main

import "fmt"

// The function 'value' returns two integers.
// The return types are specified as (int, int).
func value() (int, int) {
	return 3, 7 // returning two values at once
}

func main() {
	// Multiple assignment: capturing both returned values in 'a' and 'b'
	a, b := value()
	fmt.Println(a) // prints first returned value (3)
	fmt.Println(b) // prints second returned value (7)

	// Using blank identifier (_) when we only need the second value
	_, c := value()
	fmt.Println(c) // prints 7 (ignores the first returned value)

	// If we want to store both returned values in one variable,
	// Go doesn’t allow direct assignment like x := value()
	// So we use a struct to hold both values together
	x, y := value()
	z := struct {
		first  int // field to store first value
		second int // field to store second value
	}{x, y} // initializing struct with x and y

	// Printing individual values and struct
	fmt.Println("x", x) // prints 3
	fmt.Println("y", y) // prints 7
	fmt.Println("z", z) // prints {3 7} (struct with both values)
}

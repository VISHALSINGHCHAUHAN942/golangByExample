package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func main() {
	sum := add(2, 3)
	fmt.Println(sum)

	difference := sub(3, 2)
	fmt.Println(difference)
}

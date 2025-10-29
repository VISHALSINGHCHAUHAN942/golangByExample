package main

import (
	"fmt"
	"time"
)

func main() {
	i := 3

	// Example 1: Basic switch case with an integer value
	fmt.Print("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Print("one")
	case 2:
		fmt.Print("two")
	case 3:
		fmt.Print("three")
	}

	fmt.Println() // To move to the next line after printing number text

	// Example 2: Switch based on current weekday
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// Example 3: Switch without an expression (acts like if-else chain)
	t := time.Now()
	switch {
	case t.Hour() >= 18:
		fmt.Println("It's evening or later (after 6 PM)")
	case t.Hour() < 12:
		fmt.Println("It's before 12 PM (morning)")
	default:
		fmt.Println("It's afternoon (between 12 PM and 6 PM)")
	}
}

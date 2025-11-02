package main

import "fmt"

// 🔹 Variadic Functions in Go:
// A variadic function is a function that can accept *any number* of arguments of the same type.
// Inside the function, these arguments are treated as a slice.
// Syntax: func name(param ...Type) { ... }

// Example: fmt.Println() is a variadic function — it can take any number of arguments.


// add() is a variadic function that accepts any number of integers and prints their sum.
func add(nums ...int) {
    // The parameter 'nums' is a slice ([]int) containing all the passed integers
    fmt.Println(nums) // Print all numbers received

    total := 0
    // Loop through the slice using index-based iteration
    for i := 0; i < len(nums); i++ {
        total += nums[i] // Add each element to total
    }

    // Print the calculated sum
    fmt.Println(total)
}


// sum() is another variadic function that returns the total sum instead of printing it.
func sum(arr ...int) int {
    totall := 0
    // Loop through the slice and calculate the total
    for i := 0; i < len(arr); i++ {
        totall += arr[i]
    }

    // Return the sum value
    return totall
}


func main() {
    // Create a slice of integers
    arr := []int{1, 2, 3, 4, 5, 6}

    // To pass a slice into a variadic function, use the '...' syntax to unpack it.
    fmt.Println("sum", sum(arr...))

    // Calling add() with different numbers of arguments
    add(1, 2, 3, 4)
    add(1, 2, 3, 4, 6, 7, 8)
}

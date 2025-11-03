package main

import "fmt"

// counter returns a function (closure) that increments and returns a counter value
func counter() func() int {
    count := 0 // local variable captured by the closure
    return func() int {
        count = count + 1
        return count
    }
}

func main() {
    // counter() returns a closure function
    answer := counter()

    // Each time you call answer(), it increments its own count
    fmt.Println(answer()) // 1
    fmt.Println(answer()) // 2
    fmt.Println(answer()) // 3

    // If you call counter() again, you get a NEW closure with a fresh count
    newCounter := counter()
    fmt.Println(newCounter()) // 1
}

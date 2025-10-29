package main

import "fmt"

func main() {
    // 1️⃣ While-style loop
    // --------------------
    // This is the most basic type of for loop in Go.
    // It works like a "while" loop in other languages.
    // The loop continues as long as the condition (i <= 4) is true.
    i := 1
    for i <= 4 {
        fmt.Println(i)
        i++
    }

    // 2️⃣ Classic for loop
    // --------------------
    // This is the traditional form of a for loop:
    // - initialization (j := 0)
    // - condition (j < 4)
    // - post statement (j++)
    // It runs exactly 4 times, printing 0 to 3.
    for j := 0; j < 4; j++ {
        fmt.Println(j)
    }

    // 3️⃣ Range over integer (Go 1.22+)
    // ---------------------------------
    // Introduced in Go 1.22 — you can now use `range` over an integer value.
    // This is another way to loop N times without manually managing an index.
    // Here, it loops from 0 to 2 (3 iterations total).
    for k := range 3 {
        fmt.Println("range", k)
    }

    // 4️⃣ Infinite loop with break
    // -----------------------------
    // A `for` with no condition means it will run forever,
    // unless you explicitly break out of it or return.
    // Commonly used for continuous processes (e.g., server loops).
    for {
        fmt.Println("loop")
        break // exits after one iteration in this example
    }

    // 5️⃣ Loop with continue
    // -----------------------
    // The `continue` statement skips to the next iteration of the loop.
    // Here, we only print odd numbers (skip even ones).
    for n := range 6 {
        if n%2 == 0 {
            continue // skip even numbers
        }
        fmt.Println(n)
    }
}

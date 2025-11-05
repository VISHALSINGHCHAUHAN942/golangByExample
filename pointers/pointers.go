package main

import "fmt"

// zeroval receives the value by **copy** (pass-by-value).
// So the variable `ival` inside this function is only a COPY of `i`.
// Changing `ival` does NOT change the original `i` in main.
func zeroval(ival int) {
	ival = 0
}

// zeroptr receives the **address** of `i` (pass-by-pointer).
// `iptr` is a pointer that points to the actual memory location of `i`.
// Using *iptr allows us to change the value stored at that location,
// so this function DOES modify the original `i`.

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	// Declare and initialize variable i
	i := 1
	fmt.Println("Initial value of i:", i)

	// Pass-by-value: only a copy of i is sent to zeroval
	zeroval(i)
	// i is still 1 because zeroval modifies only the copy
	fmt.Println("After zeroval (pass-by-value), i is still:", i)

	// Pass-by-pointer: we send the address of i using &i
	zeroptr(&i)
	// Now, because the function worked on the actual memory address,
	// i gets updated to 0.
	fmt.Println("After zeroptr (pass-by-pointer), i is now:", i)

	// Printing the memory address of i
	fmt.Println("Memory address of i:", &i)
}

// | Function      | How it receives `i` | What it modifies | Result                |
// | ------------- | ------------------- | ---------------- | --------------------- |
// | `zeroval(i)`  | Copy of `i`         | Only the copy    | `i` remains unchanged |
// | `zeroptr(&i)` | Address of `i`      | Actual variable  | `i` gets changed      |


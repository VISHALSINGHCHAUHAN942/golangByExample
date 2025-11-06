package main

import (
	"fmt"
)

type person struct {
	Name       string
	Age        int
	email      string
	Isverified bool
}

// newPerson is acting like a constructor function.
// It creates a *new* person struct and returns a pointer to it.
func newPerson(name string) *person {
	p := person{Name: name} // create struct
	p.Age = 42              // set default age
	return &p               // return pointer to struct
}

func main() {

	// Creating a struct by assigning values directly
	p1 := person{"Vishal", 18, "Vishal@gmail.com", true}
	fmt.Println("The value of p1 is:", p1)

	// Creating a new person using constructor-like function
	p3 := newPerson("gillu")
	fmt.Println("p3:", p3) // p3 is a *pointer*, prints &{...}

	fmt.Println()

	// Creating struct using named fields (order doesn’t matter here)
	fmt.Println(person{Name: "Tinu", Age: 42, email: "test@gmail.com", Isverified: false})

	// Assigning struct values and accessing fields
	s := person{Name: "Sean", Age: 50}
	fmt.Println("Name of s:", s.Name, "Age of s:", s.Age)

	// Showing pointer usage
	sp := &s
	fmt.Println("Age through pointer sp:", sp.Age)

	// Updating field through pointer updates the original struct
	sp.Age = 51
	fmt.Println("Age of sp:", sp.Age)
	fmt.Println("After updating age of s using sp:", s.Age)

	// Anonymous struct example
	// Sometimes, you need a struct only once — just to store some grouped data temporarily.
    // In that case, you do NOT need to define a struct type with a name.
   // Instead, you can create a struct without a name → which is called an anonymous struct.

	dog := struct {
		name   string
		isGood bool
	}{
		"rex",
		true,
	}

	fmt.Println(dog)
}

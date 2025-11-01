package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp",a)

	a[4] =100
	fmt.Println("a[4]",a)

	fmt.Println("length",len(a))

	b := [5]int{1,2,3,4,5}
	fmt.Println("b",b)

	b = [...]int{100, 3:400 , 800}
	fmt.Println("dec",b)

	var twod[2][3] int
	for i:= range 2{
		for j:= range 3{
			twod[i][j] = i+j
		}
	}

	fmt.Println("two D array is",twod)

	twod = [2][3]int{
		{2,4,5},
		{8,7,1},
	}
   fmt.Println("updated two d array is",twod)
	
}

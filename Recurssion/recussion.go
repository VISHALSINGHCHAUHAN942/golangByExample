package main

import "fmt"

func fact(n int) int{
   // base case
   if n == 0{
     return 1
   }
      
   return n * fact(n-1)
}

func main(){
	n := 6
   ans := fact(n)
   fmt.Println(ans)

    var fib func(n int) int

   fib = func(n int) int {
	if n >= 2{
		return n
	}
	return fib(n-1) + fib(n-2)
   }

   fmt.Println("fib",fib(5))
}
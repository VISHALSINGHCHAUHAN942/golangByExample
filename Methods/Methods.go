package main

import "fmt"

type rect struct{
    length , breadth int
}

func(r * rect) area() int{
     area := r.length * r.breadth
     return area
}

func (r * rect) changeValue(){
    r.length = 25
    r.breadth = 46
}

func (r * rect) parameter() int{
    perameter := 2 * (r.length + r.breadth)
    return perameter
}

func main(){
   r := rect{length:21,breadth:22}
   fmt.Println("The area of r is",r.area())
   fmt.Println("parameter of rectangle is", r.parameter())
   r.changeValue()
    fmt.Println("length is", r.length , "Breadth is",r.breadth)
}
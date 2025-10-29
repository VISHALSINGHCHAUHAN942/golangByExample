package main

import (
	"fmt"
	"time"
)

func main(){
	i := 3

	fmt.Print("write ",i ,"as ")

	switch i{
      case 1:{
        fmt.Print("one")
	  }
	   case 2:{
        fmt.Print("two")
	  }
	case 3:{
		fmt.Print("three")
	}
	}

	  switch time.Now().Weekday(){
	  case time.Saturday , time.Sunday :
			fmt.Println("It's the weekend")
	  default:
		fmt.Println("its weekday")
	  }


      t :=time.Now()
	  switch {
	  case t.Hour() >= 18:
		   fmt.Println("its more than 12 pm")
	case t.Hour() < 12:
		fmt.Println("its before 12 pm")
	}
}
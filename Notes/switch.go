package main

import (
	"fmt"
)

func main() {
	// simple switch : used for multiple conditions
	// no need to use break in go
	// i := 5
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("other")
	// }

	// multiple considition switch
	// switch time.Now().Weekday() {
	// case time.Saturday , time.Sunday: //here we have 2 coditions
	// 	fmt.Println("weekend")
	// default:
	// 	fmt.Println("its workday")
	// }

	// type switch
	whoami := func(i interface{}) { //interface mean any type of value can be recieved
		switch t := i.(type) {
		case int:
			fmt.Println("integer")

		case string:
			fmt.Println("Its a string")
		case bool:
			fmt.Println("Its a boolean")
		default:
			fmt.Println("other", t)
		}

	}
	whoami(3)

}

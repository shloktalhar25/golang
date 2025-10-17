package main

import "fmt"

func main() {
	//if-else
	// age := 16
	// if age >= 18 {
	// 	fmt.Println("person is adult")
	// } else {
	// 	fmt.Println("is not adult")

	// }

	//else-if
	// age := 12

	// if age >= 18 {
	// 	fmt.Println("ADult")
	// } else if age >= 12 {
	// 	fmt.Println("teenager")
	// } else {
	// 	fmt.Println("kid")
	// }



	// logical operators

	// var role = "admin"
	// var hasPermissions = false

	// if role =="admin" || hasPermissions {
	// 	fmt.Println("yes")
	// }

	// if role =="admin" && hasPermissions {
	// 	fmt.Println("yes")
	// }


	// declaring variable inside if construct

	if age := 13; age >=18{
		fmt.Println("person is adult",age)
	} else if age >=12 {
		fmt.Println("teenager age is",age)
	}

	//go doesnt have ternary operator, you will have to use normal if else

}

package main

import "fmt"

const age = 20 //possible to declare outside the fucntion
// name :="golang" //not possible to use := outside the main function
func main() {
	const name = "golang"
	// name = js (not possible)
	fmt.Println(age)

	// constant grouping (multiple costant declaring)
	const (
		port = 5000
		host = "localhost"
	)
	//host =  3300 not possible to change / modify const
	fmt.Println(port, host)
}

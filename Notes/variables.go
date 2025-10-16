package main

import "fmt"

func main() {
	//declaring string variable(tip : if variable declared then it must be used)
	var name string = "golang"
	fmt.Println(name)

	//shortning
	// infer the type automatically without mentionning sting(any datatype)
	var name1 = "another golang"
	fmt.Println(name1)

	//booleen
	var is_adult= true // automatically detects booleen
	fmt.Println(is_adult)

	var age int = 34
	fmt.Println(age)

	// shorthand use / syntax
	// declare + assign 
	name2 :="golang"
	fmt.Println(name2)

	// for later assining values

	var name4 string
	name4 = "golang"
	fmt.Println(name4)

	var price float32 = 50.5
	fmt.Println(price)

	var price2 = 50.2
	fmt.Println(price2)

	price3 := 50.4
	fmt.Println(price3)



}


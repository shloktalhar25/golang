package main

import "fmt"

// maps -> hash , object , dict
func main() {
	//creating map
	
	m:= make(map[string]string)

	//setting an element 
	m["name"] = "golang"

	//get an element

	fmt.Println(m["name"])
	
	// adding element
	m["area"] = "backend"

	fmt.Println(m["name"],m["area"])

	//IMP: if key does not exist in the map it returns zero value
	
	n:= make(map[string]int)
	n["age"] = 30
	fmt.Println(n["Phone"])


	//getting length of map

	fmt.Println(len(n))


	//deleting elemnt from map

	delete(m,"price")
	fmt.Println(m)

	//clear fucntion
	clear(m)
	fmt.Println(m)

	//making a map another method

	z:= map[string]int{"price":40,"phone":3}
	fmt.Println(z)

	//
	_,ok := m["price"]

	if ok {
		fmt.Println("all ok")
	}else {
		fmt.Println("not ok")
	}


	//Map-> maps package contain equal method to compare 2 maps

}	




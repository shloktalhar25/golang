package main

import "fmt"

// numbered squence of specific length
func main() {
	var nums[10]int

	//array length
	fmt.Println(len(nums))

	//adding a elements into array
	nums[0] = 1
	fmt.Println(nums)
	
	// boolean array
	var vals[5]bool
	fmt.Println(vals) // by default all false value

	// assigning values

	vals[2] = true
	fmt.Println(vals)

	// strings

	var name[3]string
	fmt.Println(name)

	//assignging values

	name[0]= "kim" 
	fmt.Println(name)

	// int --> 0 , bool -> false , --> string--> ''

	// array delclaration + adding elements
	number:=[3]int{1,2,3}
	fmt.Println(number)

	//2-d arrays
	num2:=[2][2]int{{1,2},{3,4}}
	fmt.Println(num2)


	//fixed size ,that is predictable
	// memory optimization
	// constant time access
	// slices are used more in go than arrays
}



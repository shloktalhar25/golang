package main

import "fmt"


func sum(nums ...int)int {
	total := 0
	for _,num:= range nums {
		total = total +num
	}

	return total
}



func main() {
	fmt.Println(1,2,3,4 ,"Hello")


	 
	// result := sum(1,2,4,5,53,2,2)
	// fmt.Println(result)


	//slice
	nums := []int{3,4,5,6}
	result := sum(nums...)
	fmt.Println(result)
}
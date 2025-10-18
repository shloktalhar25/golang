package main

import "fmt"


//by value  
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("in changenum",num)
// }



//passing by reference
//dereference
func changeNum(num *int) {
	*num = 4
	fmt.Println("in changenum",*num)
}



// pointers are variables ki memory location ka address
func main() {


	//num:= 3
// 	changeNum(num)
// 	fmt.Println("change num in main",num)


	//use '&num' and '*' for  pointers
	//pointers refrers to the memory location and then making change in the real variable and not in the copy of it
	//fmt.Println("memory address",&num)


	num:= 6
	changeNum(&num)
	
	fmt.Println("memory address",num)
	
} 
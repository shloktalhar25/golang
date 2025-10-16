package main

import "fmt"

// only for loop is used inside go,no while
// we can use for to implement while
// for is only construct in go for looping
func main(){
	// while-loop implementation
	// i := 1
	// for i<=3 {
	// 	fmt.Println(i)
	// 	i =i+1
	// }

	//  infinite loop 
	// for {
	// 	fmt.Println("1")
	// }

	//classic for loop
	// for i :=0 ; i <=3 ; i++{
	// 	//break
	// 	//continue : skips the current iteration
	// 	if i == 2{
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }


	//range 

	for i := range 10 {
		fmt.Println(i)
	}
}
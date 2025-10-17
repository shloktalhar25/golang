package main

import (
	"fmt"
	"slices"
)

//slice - > dynamic array , no need to give length as we gave in array
// most used construct in go
// + usefull methods
func main() {
	//uninitialised slices is nill 
	var nums[]int

	fmt.Println(nums== nil)

	fmt.Println(len(nums))

	var nums1 = make([]int,2,5)
	fmt.Println(cap(nums1)) // prints the capacity(max no of elements that can fit)
	fmt.Println(nums1)

	//adding elemtnts usign append
	nums1= append(nums1,2,3)
	fmt.Println(nums1)
	fmt.Println(cap(nums1)) //capacity of num1

	// capacity increases automatically
	// capacity got doubled
	nums1= append(nums1,2,3,5,6)
	fmt.Println(nums1)
	fmt.Println(cap(nums1))

	//indexing
	nums1[0] =2
	fmt.Println(nums)

	// copying slice


	var nos= make([]int,2,5)
	var nos2 = make([]int,len(nos))

	nos = append(nos,2)
	fmt.Println(nos,nos2)

	copy(nos2,nos)


	// slice operator

	var nums4 = []int{1,2,3}
	fmt.Println(nums4[0:2]) // slicing just like python


	//slice -->package
	//equal from slice package
	var nums5 = []int{1,32}
	var nums6 = []int{1,2}
	fmt.Println(slices.Equal(nums5,nums6))

	// 2-d Slices

	var a = [][]int{{1,2,3},{4,5,6}}
	fmt.Println(a)

}
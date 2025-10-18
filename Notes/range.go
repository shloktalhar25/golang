package main

import "fmt"

//range is used for iteration over data structures
func main() {
	nums := []int{6, 7, 8}
	//usimg for lloop
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])

	}

	//using range to make sum
	//' _,' or 'i' is index accesing
	sum := 0
	for i, num := range nums {
		fmt.Println(i)
		sum = sum + num

	}
	fmt.Println(sum)

	//map

	m := map[string]string{"fname": "john", "lname": "doe"}
	for k, v := range m {
		fmt.Println(k, v)
	}

	//for accessing only thr keys or only the values

	for k := range m {
		fmt.Println(k)
	}

	// using range over string
	//unicode point rune
	// i is starting byte of rune
	for i, c := range "golang" {
		fmt.Println(i, string(c)) //gives unicode of word if we dont use string(c)
	}

}

package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

//returning multipe values in go

func languages() (string, string, bool) {
	return "golang", "js", true
}

// //recive function
// func processIt(fn func(a int) int) {
// 	fn(1)
// }

func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

// entry point function
func main() {
	result := add(3, 5)
	fmt.Println(result)

	//calling fucntion
	fmt.Println(languages())
	lang1, lang2, lang3 := languages()
	//or
	fmt.Println(lang1, lang2, lang3)

	//not using any value ex printing only 2 leaving while leave the 3rd as _

	fmt.Println(languages())
	language1, language2, _ := languages()
	//or
	fmt.Println(language1, language2)

	//calling prcessIt
// 	fn := func(a int) int {
// 		return 2
// 	}
// 	fn := processIt()
// 	fn(2)
// 
}

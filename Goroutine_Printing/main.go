// level 1 : Basic concepts
// part1:goroutine basics: printing concurrently

//Learn how goroutines run functions independently.
//Example: Print “Hello” and “World” concurrently.




package main

import (
	"fmt"
	"time"
)


// printNumber is a simple function that print number from 1 to 5

func printNumber() {
	for i:=1;i<=5;i++{
		fmt.Println("Number:",i)
		time.Sleep(500 * time.Millisecond) //simulate some delay
	}
}



// function to print letter from A to E
func printLetters() {
	for ch:='A';ch<='E'; ch++ {
		fmt.Println("letter:",string(ch))
		time.Sleep(500 * time.Millisecond)
	}
}



func main() {
	//start printNumber function as a seperate goroutine
	go printNumber()

	//start printLetters function as a seperate goroutine
	go printLetters()


	// if we dont wait , the program might end before goroutine finishes.
	// so we wait for 3 seconds (enough time for both goutines to complete

	time.Sleep(3* time.Second)
	fmt.Println("main function finished")
}
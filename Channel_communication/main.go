// Channels are like pipes that connect goroutines.
// One goroutine sends data into a channel, another goroutine receives it.

// They help synchronize execution and safely share data (no need for locks or shared memory).



package main 

import (
	"fmt"
	"time"	

)

// worker is a function that simulates doing some work
// It sends the result (a string) back through the provided channel

func worker(ch chan string) {
	fmt.Println("Worker:started doing some work")

	//simulate work with 2 seconds delay
	time.Sleep(2 * time.Second)

	//send message into the channel

	ch <- "work compelted successfully"

	fmt.Println("worker:message sent to channel")

}


func main() {
	//create a channel that carries string message

	ch:= make (chan string)

	//start the worker function as a goroutine and pass the chaennel to it
	go worker(ch)

	fmt.Println("Main: Waiting for result from worker...")

	//recive the message from the channel(this line blocks until data is available)
	msg := <- ch

	// once the data is recived , execution continues

	fmt.Println("main : received message :",msg)

	fmt.Println("main:program finished")


}
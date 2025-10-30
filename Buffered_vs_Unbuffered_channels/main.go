// by default the channels are unbuffered 


package main 

import (
	"fmt"
	"time"
)

//sendMessage sends 3 message into the provided channel
func sendMessage(ch chan string){
	fmt.Println("sender: Starting to send message..")

	for i := 1; i <= 3; i++ {
		msg := fmt.Sprintf("Message %d", i)
		fmt.Println("Sender: sending ->", msg)
		ch <- msg // send message into channel (may block)
		fmt.Println("Sender: sent ->", msg)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("sender:Done sending message.")



}

func main() {
	// example 1 : Unbuffered channel
	fmt.Println("---unbuffered channel")
	unbuf :=make(chan string)

	//start a sender goroutine
	go sendMessage(unbuf)

	// Recieve message one by one

	for i := 1; i <= 3; i++ {
		msg := <-unbuf // blocks until sender sends
		fmt.Println("Receiver: received ->", msg)
	}

	time.Sleep(1 * time.Second)
	fmt.Println()

		// --- Example 2: Buffered channel ---
	fmt.Println("---- BUFFERED CHANNEL ----")

	// This channel can hold 2 messages without blocking
	buf := make(chan string, 2)

	// Start a goroutine to send messages into buffered channel
	go sendMessage(buf)

	// Simulate slow receiver (delayed reading)
	time.Sleep(2 * time.Second)

	for i := 1; i <= 3; i++ {
		msg := <-buf
		fmt.Println("Receiver: received ->", msg)
	}

	fmt.Println("Main: program finished")
}


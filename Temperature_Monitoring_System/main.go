// Module 1: The Fundamentals (why and how)

// 1. Concurrency vs. Parallelism
// Concurrency is about dealing with multiple things at once.
// → It’s structuring your program so it can handle multiple tasks in overlapping time periods, even on a single CPU core.
// → Think of a single chef juggling multiple orders: chopping, stirring, checking the oven—switching between tasks.

// Parallelism is about doing multiple things at the same time.
// → Requires multiple CPU cores.
// → Like having multiple chefs, each working on a different dish simultaneously.

// Go is designed for concurrency first. It can achieve parallelism when running on multiple cores,
// but its real strength is writing clean, concurrent code that scales well—whether on 1 core or 100.

// project :Project / Mini-Task: Write a program that launches 100 goroutines. Each goroutine should print a number from 1 to 10. Observe how the output is interleaved, demonstrating the scheduler at work.
// Then, modify it so each goroutine prints its own "ID" (0-99) 10 times, to see the non-deterministic execution more clearly.

// package main

// import "fmt"

// func main() {

// 	ch:= make(chan string)

// 	go func(){
// 		ch <- "hello from goroutine"//sending message to channel
// 	}()

// 	//receiveing mesage from go routine

// 	msgage:= <-ch
// 	fmt.Println(msgage)
// }

// Unbuffered channel (blocking nature)

//Sending (ch <- value) blocks until another goroutine receives.
//Receiving (<-ch) blocks until another goroutine sends.

// package main
// import "fmt"

// func main() {

// 	ch:= make(chan int)

// 	go func() {
// 		fmt.Println("sending value....")
// 		ch <- 10 //sending message into the channel , //blocks untils main recieves
// 		fmt.Println("value sent")

// 	}()

// 	fmt.Println("reciecvving value...")
// 	value := <-ch //blocks until goroutine sends
// 	fmt.Println(value,"recived from channel")

// }

/// buffernd channel
//A buffered channel has a fixed capacity.
//You can send multiple values into it without blocking (up to its limit).

// package main

// import "fmt"

// func main() {
//     ch := make(chan string, 2)

//     ch <- "Message 1"
//     ch <- "Message 2"
//     // ch <- "Message 3" // would block if uncommented (buffer full)

//     fmt.Println(<-ch)
//     fmt.Println(<-ch)
// }

// channel Direction

//You can restrict channel direction to
// send-only or receive-only in function parameters.

//chan<- int: only send allowed
//<-chan int: only receive allowed
//Useful for safety and clarity in concurrent APIs.

// package main

// import "fmt"
// func sender(ch chan <- int ) { //send only
// 	ch <- 12
// }

// func receiver(ch <- chan int){ //rececive only
// 	fmt.Print(<-ch)

// }

// func main() {
// 	ch:=make(chan int)
// 	go sender(ch)
// 	go receiver(ch)
// }

//select statement
//select lets a goroutine wait on multiple channel operations.
//Think of it like a switch but for channels.

// package main

// import "fmt"

// func main() {
// 	select {
// 	case msg1 := <-ch1:
// 		fmt.Println("recived:",msg1)
// 	case msg2 := <-ch2:
// 		fmt.Println("recived:",msg2)
// 	default:
// 		fmt.Println("no message yet")
// 	}
// }

// Multiplexing with select
//You can receive from multiple channels simultaneously —
//whichever is ready first is chosen.

// package main

// import(
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch1:= make(chan string)
// 	ch2:= make(chan string)

// 	//this annonemous func does not run instantly
// 	go func(){
// 		time.Sleep(1 * time.Second)
// 		ch1 <-"from channel 1"
// 	}()

// 	go func() {
//         time.Sleep(2 * time.Second)
//         ch2 <- "from channel 2"
//     }()

// 	 for i := 0; i < 2; i++ {
//         select {
//         case msg1 := <-ch1:
//             fmt.Println(msg1)
//         case msg2 := <-ch2:
//             fmt.Println(msg2)
//         }
//     }
// }

// Using select for Timeouts (with time.After)

//You can use time.After to create a timeout channel.
// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     ch := make(chan string)

//     go func() {
//         time.Sleep(3 * time.Second)
//         ch <- "response"
//     }()

//     select {
//     case msg := <-ch:
//         fmt.Println("Received:", msg)
//     case <-time.After(2 * time.Second):
//         fmt.Println("Timeout! No response.")
//     }
// }

// Using select for Cancellation (done channel pattern)

// A done channel is used to signal cancellation or stop execution.

// package main

// import (
//     "fmt"
//     "time"
// )

// func worker(done <-chan bool) {
//     for {
//         select {
//         case <-done:
//             fmt.Println("Worker stopped.")
//             return
//         default:
//             fmt.Println("Working...")
//             time.Sleep(500 * time.Millisecond)
//         }
//     }
// }

// func main() {
//     done := make(chan bool)
//     go worker(done)

//     time.Sleep(2 * time.Second)
//     done <- true // signal to stop
//     time.Sleep(1 * time.Second)
// }

// 10. The close Keyword

// Use close(ch) to close a channel when no more values will be sent.

// Receivers can still read remaining values but cannot send new ones.

// Example:
// package main

// import "fmt"

// func main() {
//     ch := make(chan int, 3)
//     ch <- 1
//     ch <- 2
//     close(ch)

//     // ch <- 3 //  panic: send on closed channel

//     for val := range ch {
//         fmt.Println(val)
//     }
// }

// The range Keyword on Channels

// You can iterate over a channel using for range, which automatically stops when the channel is closed.

// Example:
// package main

// import "fmt"

// func main() {
//     ch := make(chan int)

//     go func() {
//         for i := 1; i <= 5; i++ {
//             ch <- i
//         }
//         close(ch)
//     }()

//     for num := range ch {
//         fmt.Println(num)
//     }
// }

package main

func main() {

}

// Concept	Description	Blocking?
// Unbuffered Channel	No capacity; syncs send/receive	Yes
// Buffered Channel	Has capacity; allows async send	Until buffer full
// Directional Channel	Send-only or receive-only	Type-safe
// select	Wait on multiple channels	No
// close	Ends channel communication	N/A
// range on channel	Iterates until channel closed	Auto-stop

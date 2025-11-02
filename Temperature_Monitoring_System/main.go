// Scenario

// Imagine you’re building a simple system that monitors temperature readings from a single sensor.
// The sensor sends temperature readings to a channel every few milliseconds.
// The main process:

// Reads data through a channel

// Times out if the sensor doesn’t respond

// Can be stopped with a cancel signal

// Gracefully closes the channel once done

//This is realistic — similar to how IoT or monitoring services communicate asynchronously between devices and main controllers.

package main

import (
	"fmt"
	"math/rand" // to simulate random temperature reading
	"time"      // handle sleep , timeouts and delays
)

//simulateTemperatureSensor simulates a realwordl temerateure sensor
// it continously send rarndom temerature reading to the channel
// until it recieves a stop signal through the done channel.

func simulateTemperatureSensor(tempchan chan<- float64, done <-chan bool) {
	for {
		select {
		//case 1: Stop signal recieve from main groutine
		case <-done:
			fmt.Println("sensor stopped by main goroutine.")
			close(tempchan) // closes the temperature channel before exiting
			return

		//case 2 : Default path - generate and send temperature
		default:
			//simulate a temperature reading between 20 to 30 degrees
			temp := 20 + rand.Float64()*10

			//send temperaature to the channel
			// if the channel buffer is full , this send will block
			tempchan <- temp

			// wait for a random time(0-1 seconds) before sending the next reading

			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		}
	}
}

func main() {

	//seed random number generator to produce different value on each run
	rand.Seed(time.Now().UnixNano())

	// create a buffered channel to hold temeperature readings.
	//buffered channels let the sender push several readings before blocking
	tempchan := make(chan float64, 5)

	// create an unbuffered "done" channel for cancellation signalling
	done := make(chan bool)

	//Launch the sensor simulation in a seperate goroutine
	go simulateTemperatureSensor(tempchan, done)

	//set a monitoring timeout duration
	// if no temperature is recived within 5 seconds ,the ssystem times out
	timeout := time.After(5 * time.Second)

	fmt.Println("temperature monitor started..")
	fmt.Println("----------------------------")

	//continously monitor reading until timeout or stop signals
	for {
		select {
		//case1: recive a temepeature reading from the sensor
		case temp, ok := <-tempchan: // here we get 2 params , temp and a bool
			if !ok {
				//channe closed -- no more reading
				fmt.Println("Sensor channel closed. Monitoring stopped.")
				return
			}
			//else
			fmt.Printf("Recieved temeperature: %.2f*C\n", temp)

		//case2: stop the monitoring afer a timeout period
		case <-timeout:
			fmt.Println("Timeout! no readings for too long.")
			fmt.Println("Sending stop signal to sensor.")
			done <- true //send signal to stop the sensor goroutine

			// wait briefly to let the sensor clean up and close the channel
			time.Sleep(500 * time.Millisecond)
			return

			//case 3: If no data recieved for 2 seconds , warn the user
		case <-time.After(2 * time.Second):
			fmt.Println("Warning : sensor seems slow.. waiting for next reading.")
		}
	}
}

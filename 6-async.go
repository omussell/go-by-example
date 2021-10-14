// Goroutines
// Channels
// Channel Buffering
// Channel Synchronization

package main

import (
	"fmt"
	"time"
)

func main() {

	// Goroutines
	/*
	   Lightweight thread of execution
	*/

	f("direct") // Run the function as usual, synchronously

	go f("goroutine") // Invoke as a goroutine.

	go func(msg string) {
		fmt.Println(msg)
	}("going") // Can also start a goroutine for an anonymous function

	time.Sleep(time.Second) // The two functions calls are running asynchronously in separate goroutines now.
	fmt.Println("done")

	/*
	   Prints:
	   direct : 0
	   direct : 1
	   direct : 2
	   goroutine : 0
	   going
	   goroutine : 1
	   goroutine : 2
	   done
	*/

	// Channels
	/*
	   Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.
	*/

	messages := make(chan string) // Create a new channel with make(chan val-type). Channels are typed by the values they convey.

	go func() { messages <- "ping" }() // Send a value into a channel using the channel <- syntax.

	msg := <-messages // The <-channel syntax receives a value from the channel.
	fmt.Println(msg)  // Prints ping

	/*
	   Sends and receives block until both the sender and receiver are ready. This allows us to wait at the end of the program for the message without having to use any other synchronisation.
	*/

	// Channel Buffering
	/*
	   By default channels are unbuffered, they will only accept sends (chan <-) if there is a corresponding receive (<-chan) ready to receive the sent value. Buffered channels accept a limited number of values without a corresponding receiver for those values.
	*/

	mssgs := make(chan string, 2) // Make a channel of strings buffering up to 2 values

	mssgs <- "buffered" // This channel is buffered so we can send these values into a channel without the corresponding concurrent receieve
	mssgs <- "channel"

	fmt.Println(<-mssgs) // Receive the values
	fmt.Println(<-mssgs)

	/*
	   Prints:
	   buffered
	   channel
	*/

	// Channel Synchronization
	/*
	   We can use channels to synchronize execution across goroutines.
	*/

	done := make(chan bool, 1)
	go worker(done) // Start a worker goroutine, giving it the channel to notify on.

	<-done // Block until we receive a notification from the worker on the channel
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func worker(done chan bool) {
	/*
	   This is the function we'll run in a goroutine. The done channel will be used to notify another goroutine that this functions work is done.
	*/
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

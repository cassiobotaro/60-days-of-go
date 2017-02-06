package main

import (
	"fmt"
	"time"
)

// Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.
// Do not communicate by sharing memory; instead, share memory by communicating.

// SomeComputation When done send true through a channel received as parameter
func SomeComputation(done chan bool) {
	fmt.Print("working...")
	// 1 second sleep
	time.Sleep(time.Second)
	fmt.Println("done")
	// send done and close channel
	done <- true
}

// Hello receive some string through input channel, concatenate with "Hello " and returns
// a new string through output channel
// input is specified as read-only channel and output as write-only
func Hello(input <-chan string, output chan<- string) {
	received := <-input
	received = "Hello " + received
	output <- received
}

// Echo returns what receive through two way channel
func Echo(inout chan string) {
	inout <- <-inout
}

func main() {
	// syntax note
	// <- channel   receive from channel
	// channel <-   send to channel

	// sending some message back from a goroutine
	message := make(chan string)
	go func(m chan string) {
		m <- "some message"
	}(message)
	fmt.Println(<-message)

	// buffered channels are like convetional, but can store
	// more values until you request the next
	msgs := make(chan string, 2)
	msgs <- "buffered"
	msgs <- "channel"
	fmt.Println(<-msgs)
	fmt.Println(<-msgs)

	done := make(chan bool)
	go SomeComputation(done)
	// wait until some computation are done
	<-done

	// send from "in" channel and receive from "out"
	in := make(chan string)
	out := make(chan string)
	go Hello(in, out)
	in <- "John"
	fmt.Println(<-out)

	// for default channels are two way
	// we can send and receive data
	echoChan := make(chan string)
	go Echo(echoChan)
	echoChan <- "Ping"
	fmt.Println(<-echoChan)
}

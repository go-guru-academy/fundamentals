package main

import "fmt"

func main() {

	// Channels are a Go concurrency primitive.
	// Channels allow for values to be sent out of and between go routines.

	// Channels must be "made". The default value of a channel is nil, therefore,
	// if a channel is not made then no channel operations can be performed on it.
	// Channels are typed by the values they convey. Therefore, the below channel
	// is an unbuffered, "string" channel
	c1 := make(chan string)

	// Unbufferred channels block when sending or receiving unless there is
	// a party on the other side.
	// If, given our current code, we were to write the following line...
	//
	// c1 <- "hello world"
	//
	// our program would panic and say "all go routines are asleep - deadlock!".
	// If, given our current code, we were to write the following line...
	//
	// <-c1
	//
	// our program would panic and say "all go routine are asleep - deadlock!"
	// Go routines are the solution to this problem.
	//
	// Option 1: We send data to the channel within the go routine and then
	// receive data from the channel after the go routine
	go func() {
		// Sending data into the channel
		c1 <- "hello world"
	}()

	// Receiving data from the channel
	fmt.Println(<-c1)

	// Option 2: We receive data from the channel within the go routine and then
	// send data to the channel after the go routine
	go func() {
		// Receiving data from the channel
		fmt.Println(<-c1)
	}()

	// Sending data into the channel
	c1 <- "hello world"

	// Sending data to and receiving data from a channel always uses the arrow
	// operator facing the same direction <-
	// The arrow never faces the other direction
	// <-c1 receiving data from the channel
	// c1 <- "hello world" sending data to the channel

}

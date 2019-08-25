package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}

// Channels - pipes that connect concorrent goroutines
// - can send values into channels from one goroutine and
//  receive those values into another goroutine

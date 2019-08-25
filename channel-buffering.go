package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

// default channels are unbuffered - only accept sends (chan <-)
// if there is a corresponding receive (<- chan)
//! buffered channels accept a limited number of vals
//! w/o a corresponding receiver for those values

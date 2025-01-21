package main

import "fmt"

func main() {
	// Create a buffered channel with capacity 2
	bufferedChannel := make(chan string, 2)

	// Send data into the channel
	bufferedChannel <- "Message 1"
	bufferedChannel <- "Message 2"

	// Try to receive data
	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)
}

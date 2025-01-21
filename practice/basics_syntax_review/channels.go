package main

import "fmt"

func worker(channel chan string) {
	channel <- "Hello from worker!"
}

func main() {
	// Create a channel
	messageChannel := make(chan string)

	// Start a goroutine
	go worker(messageChannel)

	// Receive data from the channel
	message := <-messageChannel
	fmt.Println("Received:", message)
}

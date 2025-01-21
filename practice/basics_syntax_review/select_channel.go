package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	// Start goroutines to send data
	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "Message from Channel 1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- "Message from Channel 2"
	}()

	// Use select to wait on both channels
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Println("Received:", msg1)
		case msg2 := <-channel2:
			fmt.Println("Received:", msg2)
		}
	}
}

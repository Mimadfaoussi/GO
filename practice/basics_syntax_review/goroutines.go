package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// Start a goroutine
	go printMessage("Goroutine")

	// Run another function in the main thread
	printMessage("Main Thread")
}

package main

import (
	"fmt"
	"time"
)

func worker(channel chan string, id int) {
	time.Sleep(time.Second)
	channel <- fmt.Sprintf("Worker %d completed", id)
}

func main() {
	// Create a buffered channel
	jobChannel := make(chan string, 2)

	// Start multiple workers
	for i := 1; i <= 3; i++ {
		go worker(jobChannel, i)
	}

	// Receive results from the channel
	for i := 1; i <= 3; i++ {
		fmt.Println(<-jobChannel)
	}
}

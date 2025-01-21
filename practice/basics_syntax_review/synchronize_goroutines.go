package main

import (
	"fmt"
	"sync"
)

func printMessage(msg string, wg *sync.WaitGroup) {
	defer wg.Done() // Notify WaitGroup that this goroutine is done
	for i := 0; i < 5; i++ {
		fmt.Println(msg, i)
	}
}

func main() {
	var wg sync.WaitGroup

	// Add 1 to WaitGroup counter
	wg.Add(1)

	// Start a goroutine
	go printMessage("Goroutine", &wg)

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All goroutines completed.")
}

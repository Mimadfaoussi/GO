package main

import (
	"fmt"
	"strings"
)

func word_frequency(str string) map[string]int {
	word_freq := map[string]int{}
	var exists bool
	lowercase := strings.ToLower(str)
	words := strings.Fields(lowercase)
	for _, word := range words {
		_, exists = word_freq[word]
		if (exists) {
			word_freq[word]++
		} else {
			word_freq[word] = 1
		}
	}
	return word_freq
}

// Call the function
func main() {
	s := "This is a test. This is only a test."
	fmt.Println(word_frequency(s)) // Should output: ["this", "is", "the", "end", ":)"]
}

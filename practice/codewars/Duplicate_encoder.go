package main

import (
	"fmt"
	"strings"
)


func is_double(pos int, word string) bool {
	if (pos < 0 || pos >= len(word)) {
		fmt.Println("out of range")
		return false
	}

	lowerword := strings.ToLower(word)
	for i:=0; i < len(word); i++ {
		if (i == pos){
			continue
		}
		if (lowerword[i] == lowerword[pos]){
			return true
		}
	}
	return false
}



func DuplicateEncode(word string) string {
	new_word := ""
	for i:= range word {
		if (is_double(i, word) == true){
			new_word = new_word + ")"
		} else {
			new_word = new_word + "("
		}
	}
	return new_word
}


func main() {
	new_word := DuplicateEncode("this is a string")
	fmt.Println(new_word)
}

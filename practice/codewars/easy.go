package main

import (
	"fmt"
	"strings"
)

func reverse_string(str string) string {

	newstr := ""
	for _, chr := range str {
		newstr = string(chr) + newstr
	}
	//fmt.Println(newstr)
	return newstr
}

func palindrome(str string) bool {
	newstr := reverse_string(str)
	if (newstr == str){
		return true
	}
	return false
}

func sumArray(arr []int) int {
	sum := 0
	for _, elm := range arr {
		sum = sum + elm
	}
	fmt.Println(sum)
	return sum
}

func countVowels(str string) int {
	sum := 0
	for _, chr := range str {
		if (strings.Contains("aeoui", string(chr))) {
			sum++
		}
	}
	fmt.Printf("there are %d vowels in %s\n", sum, str)
	return sum
}


func main() {
	arr := []int{1,2,2,4,1}
	test := "golang"
	p1 := "madam"
	p2 := "not"
	p3 := "weeeew"
	fmt.Println(reverse_string(test))
	fmt.Println(palindrome(p1))
	fmt.Println(palindrome(p2))
	fmt.Println(palindrome(p3))
	sumArray(arr)
	countVowels(p1)
}

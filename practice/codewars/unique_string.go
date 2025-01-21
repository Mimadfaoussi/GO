package main

import (
	"fmt"
	"strings"
)

func unique(str string) bool {
	lower := strings.ToLower(str)
	done := ""
	for char:= range(lower){
		if strings.Contains(done, string(str[char])) == true {
			return false
		}
		done = done + string(str[char])
	}
	return true
}

func test_cases() {
	test1 := "this"
	test2 := "correct"
	test3 := "No"
	test4 := "JAJA"
	fmt.Printf("%s -> %v\n",test1,unique(test1))
	fmt.Printf("%s -> %v\n",test2,unique(test2))
	fmt.Printf("%s -> %v\n",test3,unique(test3))
	fmt.Printf("%s -> %v\n",test4,unique(test4))

}

func main(){
	test_cases()
	//fmt.Println(unique("this"))
}

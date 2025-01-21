package main

import (
	"fmt"
)

func is_prime(nb int) bool {
	if (nb == 1) {
		return false
	}
	if (nb == 2) {
		return true
	}
	i := 2
	for i = 2; i <= nb; i++ {
		if (nb % i == 0) {
			return false
		}
		if (i * i > nb){
			break
		}
	}
	return true
}

func main() {
	for i:=1; i <= 100; i++ {
		if (is_prime(i) == true) {
			fmt.Println("Prime")
			continue
		}
		if (i % 3 == 0 && i % 5 == 0){
			fmt.Println("FizzBuzz")
			continue
		}
		if (i % 3 == 0) {
			fmt.Println("Fizz")
		} else if (i % 5 == 0) {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

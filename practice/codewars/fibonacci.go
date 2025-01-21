package main

import "fmt"


func fibonacci(n int) int {
	ex := 0
	sum := 1
	for i:=0; i < n - 1; i++ {
		mid := sum
		sum = sum + ex
		ex = mid
	}
	return sum
}

func main() {
	fmt.Println(fibonacci(10))
}

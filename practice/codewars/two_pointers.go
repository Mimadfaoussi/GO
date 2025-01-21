package main

import (
	"fmt"
)

func two_sum(arr []int, target int) []int {
	if len(arr) <= 1 {
		return []int{}
	}
	for i := 0; i < len(arr) - 1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if (arr[i] + arr[j] == target) {
				res := []int{arr[i], arr[j]}
				return res
			}
		}
	}
	return []int{}
}

func main() {
	arr := []int{1,2,3,4,5,6,7,8,9}
	target := 8
	fmt.Println(two_sum(arr, target))
}

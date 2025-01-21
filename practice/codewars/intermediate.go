package main

import (
	"fmt"
//	"strings"
)


func findlargest(arr []int) int{
	if (len(arr) == 0) {
		fmt.Println("empty array")
		return 0
	}
	largest := arr[0]
	for _, elm := range (arr) {
		if (elm > largest) {
			largest = elm
		}
	}
	fmt.Printf("largest is : %d\n", largest)
	return largest
}

func secondLargest(arr []int) int{
	largest := findlargest(arr)
	second := arr[0]
	for _, elm := range (arr){
		if elm < largest && elm > second {
			second = elm
		}
	}
	fmt.Printf("second largest : %d\n", second)
	return second
}


func missingNumber(arr []int) int {
	if len(arr) == 0 {
		fmt.Println("no list provided")
		return 0
	}
	for i:= 0; i < len(arr) - 1; i++ {
		if (arr[i] + 1 != arr[i + 1]) {
			fmt.Printf("the missing number is : %d\n", arr[i] + 1)
			return arr[i] + 1
		}
	}
	fmt.Println("no missing number found\n")
	return -1
}


func matrix_Transpose(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return [][]int {}
	}
	rows := len(matrix)
	cols := len(matrix[0])
	transpose := make([][]int, rows)
	for i:= range transpose {
		transpose[i] = make([]int, 0, cols)
	}
	for _, ligne := range (matrix) {
		for index, elm := range (ligne) {
			fmt.Printf("index: %d, elm: %d\n", index, elm)
			transpose[index] = append(transpose[index], elm)
		}
	}
	fmt.Println(transpose)
	return transpose
}

func main() {
	matrix := [][]int {{1,2,3}, {4,5,6}, {7,8,9}}
//	arr := []int{1,2,3,5,-119,4,5,6,3}
	//arr2 := []int{1,2,3,4,5,6,7,8,9,11,12}
	//secondLargest(arr2)
	//missingNumber(arr2)
	matrix_Transpose(matrix)
}





















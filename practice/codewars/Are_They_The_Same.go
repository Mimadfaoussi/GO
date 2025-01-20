package main

import (
	"fmt"
)


func is_square(i int, array2 []int) bool{
	power2 := i * i
	for j := range array2 {
		if (power2 == array2[j]){
			return true
		}
	}
	return (false)
}


func Comp(array1 []int, array2 []int) bool {
	if (arra1 == nil || arra2 == nil) {
		return false
	}
	if (len(array1) == 0 || len(array2) == 0 || len(array1) != len(array2)){
		return false
	}
	for i := range array1 {
		if (is_square(array1[i], array2) == false) {
			return false
		}
	}
	return (true)
}


func main() {

	
        var a1 := []int{121, 144, 19, 161, 19, 144, 19, 11}
        var a2 := []int{11*11, 121*121, 144*144, 19*19, 161*161, 19*19, 144*144, 19*19}

	fmt.Println(Comp(a1, a2))
}

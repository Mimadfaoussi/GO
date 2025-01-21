package main

import "fmt"

func main() {
    // Array: fixed size
    var arr [3]int = [3]int{1, 2, 3}
    fmt.Println("Array:", arr)

    // Slice: dynamic size
    slice := []int{4, 5, 6}
    fmt.Println("Slice:", slice)

    // Append to a slice
    slice = append(slice, 7, 8)
    fmt.Println("After append:", slice)

    // Access and modify elements
    fmt.Println("First element of array:", arr[0])
    slice[1] = 50
    fmt.Println("Modified slice:", slice)

    // Loop through slices
    fmt.Println("Looping through slice:")
    for index, value := range slice {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

    // Slicing a slice
    subSlice := slice[1:4] // includes index 1, excludes index 4
    fmt.Println("Sub-slice:", subSlice)
}


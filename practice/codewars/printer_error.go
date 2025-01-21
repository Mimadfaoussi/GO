package main

import (
	"fmt"
)



func PrinterError(s string) string {
	l := len(s)
	count := 0
	for _, chr := range s{
		if (chr > 'm') {
			count++
		}
	}
	newstr := fmt.Sprintf("%v/%v", count, l)
	fmt.Println(newstr)
	return newstr
}

func main() {
	PrinterError("aaaxbbbbyyhwawiwjjjwwm")
}

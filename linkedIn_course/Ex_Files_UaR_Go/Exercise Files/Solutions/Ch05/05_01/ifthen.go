package main

import (
	"fmt"
	"strconv"
)


func main() {

	//var x float64 = 42
	var result string

	if x := -42; x < 0 {
		result = "Less than zero"
	} else if x == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println("Result:", result)
//	fmt.Println("Value of x:", x)

	days := 6
	s := strconv.Itoa(days)
	fmt.Println(s, days)
}
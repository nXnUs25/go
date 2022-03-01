package main

import "fmt"

func main() {

	// var x float64 = 0
	var result string

	if x := -1 ; x < 0 {
		result = "Less than zero"
	} else if x > 0 {
		result = "Greater than zero"
	} else {
		result = "Equal to zero"
	}

	fmt.Printf("%T: %v", result, result)
	//fmt.Printf("You cannot access the x var, it has only if scope \n%T: %v", x, x)
}

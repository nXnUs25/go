package main

import (
	"fmt"
	"mastering_go/ch01_exercises"
	"os"
)

func main() {
	// if we using os.Args first index 0 contains always file name
	// so we strats from 1: to the end
	sum, avg := ch01_exercises.CalculateAvg(os.Args[1:])
	fmt.Printf("[Result]: The Sum of %s is equal to: %.2f and the AVG is equal to %.2f \n", os.Args[1:], sum, avg)
}

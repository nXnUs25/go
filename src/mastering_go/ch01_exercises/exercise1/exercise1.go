package main

import (
	"fmt"
	"mastering_go/ch01_exercises"
	"os"
)

func main() {
	// if we using os.Args first index 0 contains always file name
	// so we strats from 1: to the end
	sumInts := ch01_exercises.CalculateSum(os.Args[1:])
	fmt.Printf("[Result]: The sum of %s is equal to: %d\n", os.Args[1:], sumInts)

	sumFloats := ch01_exercises.CalculateFloatSum(os.Args[1:])
	fmt.Printf("[Result]: The sum of %s is equal to: %f\n", os.Args[1:], sumFloats)
}

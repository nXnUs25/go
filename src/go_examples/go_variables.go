package main

import (
	"fmt"
)

func main() {
	fmt.Println("Testing declarations assignment and variables")

	/* boolean assignment */
	var logic bool = true
	fmt.Println("Variable logic:", logic)

	/* int */
	var x int
	x = 111
	fmt.Println("x: ", x)
	var y int = 23
	fmt.Println("y: ", y)
	var z = 400
	fmt.Println("z: ", z)

	/* assaigning multiple values */
	var i, j = 20, "testing this"
	fmt.Println("Value i: ", i, "Value j: ", j)
	fmt.Println("Value i and j: ", i, j)

	/* quick assign */
	a := 20
	b := 3.8
	fmt.Println("Value a: ", a, "\nValue b: ", b)

	/* changing value may have error as its declaration */
	//b := 22
	// a = 4.8
	fmt.Println("Value a: ", a, "\nValue b: ", b)

}

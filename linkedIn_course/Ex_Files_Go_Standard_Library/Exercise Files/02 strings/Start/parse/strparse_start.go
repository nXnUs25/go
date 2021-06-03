package main

import (
	"fmt"
	"strconv"
)

func main() {
	sampleint := 100
	samplestr := "250"

	// This does a character conversion, not a numerical one
	fmt.Println("\n// This does a character conversion, not a numerical one")
	newstr := string(sampleint)
	fmt.Println("Int is: :", sampleint)
	fmt.Println("Result of using string():", newstr)

	// The strconv package contains a variety of functions for parsing and formatting
	// numbers, values, and strings
	fmt.Println("\n// The strconv package contains a variety of functions for parsing and formatting")
	fmt.Println("// numbers, values, and strings")

	// Convert an integer to string
	fmt.Println("\n// Convert an integer to string")
	s := strconv.Itoa(sampleint)
	fmt.Printf("%T, %v\n", s, s)

	// Convert a string to integer
	fmt.Println("\n// Convert a string to integer")
	x, err := strconv.Atoi(samplestr)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%T, %v\n", x, x)

	// Other parse functions
	fmt.Println("\n// Other parse functions")
	b, _ := strconv.ParseBool("true")
	fmt.Println("strconv.ParseBool(\"true\")", b)
	f, _ := strconv.ParseFloat("3.14159", 64)
	fmt.Println("strconv.ParseFloat(\"3.14159\", 64)", f)
	i, _ := strconv.ParseInt("-42", 10, 64)
	fmt.Println("strconv.ParseInt(\"-42\", 10, 64)", i)
	u, _ := strconv.ParseUint("42", 10, 64)
	fmt.Println("strconv.ParseUint(\"42\", 10, 64)", u)

	// Other format functions
	fmt.Println("\n// Other format functions")
	s = strconv.FormatBool(true)
	fmt.Println("strconv.FormatBool(true)", s)
	s = strconv.FormatFloat(3.14159, 'E', -1, 64)
	fmt.Println("strconv.FormatFloat(3.14159, 'E', -1, 64)", s)
	s = strconv.FormatInt(-42, 10)
	fmt.Println("strconv.FormatInt(-42, 10)", s)
	s = strconv.FormatUint(42, 10)
	fmt.Println("strconv.FormatUint(42, 10)", s)
}

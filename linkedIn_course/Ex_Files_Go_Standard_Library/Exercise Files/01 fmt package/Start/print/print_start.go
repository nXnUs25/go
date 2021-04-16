package main

import "fmt"

func main() {
	// TODO: Print a simple string, no newline
	fmt.Print("welcome go")
	// TODO: Print with a newline
	fmt.Println("Welcome to go")

	// TODO: Print a string with values
	const answer = 42
	fmt.Println("The answer is:", answer)

	// TODO: Print a string with multiple interspersed values
	const a, b, c = 5, 5, 10
	fmt.Println("Add", a, "and", b, "equal", c)
	// TODO: print a slice of data
	items := []int{10, 20, 40, 80}

	length, err := fmt.Println(items)

	fmt.Println(length, err)

}

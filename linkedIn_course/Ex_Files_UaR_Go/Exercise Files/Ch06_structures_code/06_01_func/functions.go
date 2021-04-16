package main

import "fmt"

func main() {
	doSomething()
}

func doSomething() {
	fmt.Println("Doing something..... ")

	fmt.Println(addValues(2, 5))

	fmt.Println(sayWord("bla bla"))

	fmt.Println(addAllValues(2, 3, 5, 6))
}

func sayWord(s string) string {
	return s
}

func addValues(v1 int, v2 int) int {
	return v1 + v2
}

func addAllValues(values ...int) int {
	sum := 0

	for x := range values {
		sum += values[x]
	}

	return sum

}
package main

import "fmt"

func main() {
	var colors [3]string
	colors[0] = "red"
	colors[1] = "green"
	colors[2] = "blue"

	fmt.Println(colors)

	fmt.Println(colors[2])

	var numbers = [5]int{1,2,3,4,5}
	fmt.Println(numbers)

	l := len(colors)
	fmt.Println("Colors length ", l)
	l = len(numbers)
	fmt.Println("Numbers length ", l)
	fmt.Printf("%T, %v\n", numbers, numbers)

	capacity := cap(numbers)
	fmt.Printf("Array - %T, %v\n", capacity, capacity)
}
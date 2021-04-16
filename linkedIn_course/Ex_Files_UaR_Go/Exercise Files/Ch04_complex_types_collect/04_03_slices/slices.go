package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"red", "green", "blue"}
	fmt.Printf("%T, %v\n", colors, colors)
	colors = append(colors, "yellow")
	fmt.Printf("%T, %v\n", colors, colors)
	colors = append(colors[1:len(colors)])
	fmt.Printf("%T, %v\n", colors, colors)
	colors = append(colors[1:])
	fmt.Printf("%T, %v\n", colors, colors)
	colors = append(colors[:len(colors) - 1])
	fmt.Printf("%T, %v\n", colors, colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 122
	numbers[1] = 76789
	numbers[2] = 32
	numbers[3] = 12
	numbers[4] = 156

	fmt.Printf("%T, %v\n", numbers, numbers)

	numbers = append(numbers, 235)
	fmt.Printf("%T, %v\n", numbers, numbers)
	capacity := cap(numbers)
	fmt.Printf("Slices - %T, %v\n", capacity, capacity)

	sort.Ints(numbers)
	fmt.Printf("%T, %v\n", numbers, numbers)

}
package main

import "fmt"

// Write a Go program that converts an existing array into a map.
func main() {
	// we wanna do first that:
	// arr_value1 = key
	// arr_value2 = key
	// and so on..
	arr := [4]string{"a", "b", "c", "d"}
	conv := map[string]string{}

	for x := 0; x < len(arr); x += 2 {
		conv[arr[x]] = arr[x+1]
	}
	fmt.Printf("Map is: %v\n", conv)
	for k, v := range conv {
		fmt.Printf(" %v -> %v\n", k, v)
	}
	// converting arr into map so arr index is value of map and key is arr value
	conv2 := map[string]int{}
	for i, v := range arr {
		conv2[v] = i
	}
	fmt.Printf("Map is: %v\n", conv2)
	for k, v := range conv2 {
		fmt.Printf(" %v -> %v\n", k, v)
	}
}

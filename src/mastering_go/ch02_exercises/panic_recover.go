package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "1.2"
	v := badConverstions(a) // this way we isolated the problem just to a func moving same block from main func to separate
	fmt.Printf("Value from parse would be printed as v: %d\n", v)
	fmt.Println()
	c := badConverstions2(a) // this way we isolated the problem just to a func moving same block from main func to separate
	fmt.Printf("Value from parse would be printed as c: %d\n", c)
}

func badConverstions(a string) int {
	v, e := strconv.ParseInt(a, 10, 0)
	if e != nil {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
				fmt.Println("Recovered from Panic in if")
			}
		}()
		panic(e)
	}
	fmt.Printf("Value from parse wont be printed if panic v: %d", v)
	return int(v)
}

/// just to show that recover is always executed after panic at the end the func which caused panic
func badConverstions2(a string) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Println("Recovered from Panic in func block v2")
		}
	}()
	v, e := strconv.ParseInt(a, 10, 0)
	if e != nil {
		panic(e)
	}
	fmt.Printf("Value from parse wont be printed if panic v: %d", v)
	return int(v)
}

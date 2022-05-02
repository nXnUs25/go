package main

import (
	"fmt"
	"time"
)

// increased the loop from 0 to 100
// before was going only to 10
func function() {
	for i := 0; i < 100; i++ {
		fmt.Print(i)
	}
}

func main() {
	go function()

	// increased the loop to be from 100 to 200
	// exaple from book was 10 to 20
	// after making the loops going longer allow us to see unexpected results
	// and that was the pourpose of change
	go func() {
		for i := 100; i < 200; i++ {
			fmt.Print(i, " ")
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println()
}

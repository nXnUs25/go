package main

/*
Remove the time.Sleep(1 * time.Second) statement from the simple.go program and see what happens. Why is that?

without of this line no output will be written since func do not have time to finish ... main func finishes earlier than other 2
chan done bool works to fix this method executions.
*/

import (
	"fmt"
	//"time"
)

// increased the loop from 0 to 100
// before was going only to 10
func function(b chan bool) {
	fmt.Println("func 1")
	for i := 0; i < 100; i++ {
		fmt.Print(" f1 ", i)
	}
	fmt.Println()
	b <- true
}

func main() {
	done := make(chan bool, 2)
	go function(done)
	fmt.Println()

	// increased the loop to be from 100 to 200
	// exaple from book was 10 to 20
	// after making the loops going longer allow us to see unexpected results
	// and that was the pourpose of change
	go func(b chan bool) {
		fmt.Println("func 2")
		for i := 100; i < 200; i++ {
			fmt.Print(i, " f2: ")
		}
		fmt.Println()
		b <- true
	}(done)
	<-done

	//time.Sleep(1 * time.Second)
	fmt.Println("all done without of sleep")
}

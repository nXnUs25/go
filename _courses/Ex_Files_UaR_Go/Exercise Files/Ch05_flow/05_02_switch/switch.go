package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	dow := rand.Intn(6) + 1
	fmt.Println("Day", dow)

	result := ""

	switch dow {
		case 1:
			result = "It's Sunday"
		case 7:
			result = "It's Saturday"
		default:
			result = "It's a weekday"
	}

	fmt.Println(result)

	fmt.Println("")

	switch ran := rand.Intn(6) +1 ; ran {
		case 1:
			result = "It's Sunday"
		case 7:
			result = "It's Saturday"
			fallthrough // fall into next case
		default:
			result = "It's a weekday"
	}
	fmt.Println(result)

	x := -42;
	switch {
		case x < 0:
			result = "Less than zero"
		case x > 0:
			result = "Greater than zero"
		default:
			result = "Equal to zero"
	}
	fmt.Println(result)

/**
// You cannot do this way type bool expected and y is int
	switch y := 22 ; y{
		case y < 0:
			result = "Less than zero"
		case y > 0:
			result = "Greater than zero"
		default:
			result = "Equal to zero"
}
fmt.Println(result)
*/
}

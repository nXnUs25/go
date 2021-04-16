package main

import "fmt"

func main() {

	for x := 0; x < 6; x++ {
		fmt.Println("For loop x is now :", x)
	}

	var i int
	for i = 0; i < 6; i++ {
		fmt.Println("For loop i is now :", i)
	}

	var x = 50
	if x < 10 {
		fmt.Println("X its less than 10")
	} else {
		fmt.Println("X its greaten than 10")
	}

	if x < 10 {
		fmt.Println("X less than 10 : ", x)
	} else if x >= 10 && x <= 90 {
		fmt.Println("X greater or equal 10 and X less or equal 90: ", x)
	} else {
		fmt.Println("X grear=ter than 90 : ", x)
	}

	// testing switch case
	a, b := 3, 2

	switch a + b {
	case 2:
		fmt.Println("Value 2: ", (a + b))
	case 3:
		fmt.Println("Value 3: ", (a + b))
	case 5:
		fmt.Println("Value 5: ", (a + b))
	default:
		fmt.Println("Value its default")
	}

}

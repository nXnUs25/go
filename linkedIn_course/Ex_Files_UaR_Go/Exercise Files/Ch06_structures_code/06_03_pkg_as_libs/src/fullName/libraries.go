package main

import (
	"fmt"
	"fname"
)

func main() {

	n1, l1 := fname.FullName("Zaphod", "Beeblebrox")
	fmt.Printf("Fullname: %v, number of chars: %v\n\n", n1, l1)

	n2, l2 := fname.FullNameNakedReturn("Arthur", "Dent")
	fmt.Printf("Fullname: %v, number of chars: %v\n\n", n2, l2)

}


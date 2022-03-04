package main

import (
	"fmt"
	"mastering_go/ch03_exercises/iota_constant_generator"
)

// Can you write a constant generator iota for the powers of the number four?
func main() {
	fmt.Printf("Power of (%-5s) is %v\n", "4^4", iota_constant_generator.P4_4)
	fmt.Printf("Power of (%-5s) is %v\n", "4^5", iota_constant_generator.P4_5)
	fmt.Printf("Power of (%5s) is %v\n", "4^6", iota_constant_generator.P4_6)
	fmt.Printf("Power of (%5s) is %v\n", "4^10", iota_constant_generator.P4_10)
}

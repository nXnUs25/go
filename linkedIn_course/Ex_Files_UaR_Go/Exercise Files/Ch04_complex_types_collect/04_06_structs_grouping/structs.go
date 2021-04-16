package main

import (
	"fmt"
)

type Dog struct {
	Breed string
	Weight int
}

func main() {
	poodle := Dog{"Poodle", 32}
	fmt.Printf("%T : %v\n", poodle, poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
}
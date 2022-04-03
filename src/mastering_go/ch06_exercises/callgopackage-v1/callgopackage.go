package main

import (
	"fmt"

	v1 "github.com/nXnUs25/gopackage"
)

func main() {
	v1.ShowVersion()
	fmt.Println(v1.VERSION, "Called Version constant")
	fmt.Println(v1.Version())
}

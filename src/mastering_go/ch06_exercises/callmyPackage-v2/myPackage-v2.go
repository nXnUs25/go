package main

import (
	"fmt"

	//my "github.com/nXnUs25/go/src/mastering_go/ch06_exercises/myPackage"
	v2 "github.com/nXnUs25/myGoModule/v2"
)

func main() {
	fmt.Println(v2.PackageName())
	fmt.Println(v2.OtherFileInSamePackage())
	v2.CallPrivate("nonus25", "nonus25-lastname")
	v2.Print()
	v2.Version()
}

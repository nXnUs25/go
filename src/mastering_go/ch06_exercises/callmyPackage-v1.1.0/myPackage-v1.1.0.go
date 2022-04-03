package main

import (
	"fmt"

	//my "github.com/nXnUs25/go/src/mastering_go/ch06_exercises/myPackage"
	v11 "github.com/nXnUs25/myGoModule"
)

func main() {
	fmt.Println(v11.PackageName())
	fmt.Println(v11.OtherFileInSamePackage())
	v11.CallPrivate("nonus25")
	v11.Print()
	v11.Version()
}

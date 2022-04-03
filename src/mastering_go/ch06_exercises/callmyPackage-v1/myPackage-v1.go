package main

import (
	"fmt"

	//my "github.com/nXnUs25/go/src/mastering_go/ch06_exercises/myPackage"
	m "github.com/nXnUs25/myGoModule"
)

func main() {
	/*
		fmt.Println(my.PackageName())
		fmt.Println(my.OtherFileInSamePackage())
		my.CallPrivate("nonus25")
		my.Print()
	*/
	fmt.Println(m.PackageName())
	fmt.Println(m.OtherFileInSamePackage())
	m.CallPrivate("nonus25")
	m.Print()
	m.Version()
}

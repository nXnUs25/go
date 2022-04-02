package myPackage

import "fmt"

const version = "1.0.0"

func OtherFileInSamePackage() string {
	return "Its other file location for this file"
}

func privatePackage(name string) string {
	return name + " " + name
}

func print() {
	fmt.Println("function been added and we didn't ran `go install`")
}

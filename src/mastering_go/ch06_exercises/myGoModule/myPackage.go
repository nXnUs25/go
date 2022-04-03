package myGoModule

import "fmt"

func PackageName() string {
	return "myPackage"
}

func CallPrivate(name string) {
	fmt.Println(privatePackage(name))
}

func Version() {
	fmt.Println("Version: ", version)
}

func Print() {
	fmt.Println("Package Calling print(): ")
	print()
}

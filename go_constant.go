package main

import "fmt"

func main() {
	const b = 10

	fmt.Println("Constant VLUE B IS :", b)

	// this will generate error when uncomment
	//b = 20
	fmt.Println("Constan Value cannot be changed: b is : ", b)
}

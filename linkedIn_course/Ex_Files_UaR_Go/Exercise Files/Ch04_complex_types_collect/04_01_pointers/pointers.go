package main

import "fmt"

func main() {
	var p *int

	if p != nil{
		fmt.Println("Value of p: ", *p)
	}else{
		fmt.Println("Value of p is null")
	}

	var v int = 42
	p = &v

	if p != nil{
		fmt.Println("Value of p: ", p)
		fmt.Println("Value of p: ", *p)
	}else{
		fmt.Println("Value of p is null")
	}
	var value1 float64 = 42.13
	pointer1 := &value1

	fmt.Println("Value1 :", *pointer1)

	*pointer1 = *pointer1 / 31
	fmt.Println("Value1 :", *pointer1)
	fmt.Println("Value1 :", value1)
}

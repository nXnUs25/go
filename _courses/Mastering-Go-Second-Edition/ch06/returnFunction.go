package main

import (
	"fmt"
)

func funReturnFun() func() int {
	i := 0
	return func() int {
		i++
		fmt.Printf(">> i:[%v] -> i * i:[%v]\n", i, i*i)
		return i * i
	}
}

func main() {
	i := funReturnFun()
	j := funReturnFun()

	fmt.Println("1:", i())
	fmt.Println("2:", i())
	fmt.Println("j1:", j())
	fmt.Println("j2:", j())
	fmt.Println("3:", i())
	fmt.Println("j3:", j())
}

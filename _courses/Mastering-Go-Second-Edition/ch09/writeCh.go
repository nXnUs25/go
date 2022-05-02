package main

import (
	"fmt"
	"time"
)

func writeToChannel(c chan int, x int) {
	fmt.Println(x)
	fmt.Println(x)
	c <- x
	fmt.Println("No execution: ", x)
	close(c)
	fmt.Println("Same nothing <- blocks all code below", x)
}

func main() {
	c := make(chan int)
	go writeToChannel(c, 10)
	time.Sleep(1 * time.Second)
}

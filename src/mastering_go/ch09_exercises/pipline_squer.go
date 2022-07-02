/*
Create a pipeline for calculating the sum of the squares of all of the natural numbers in a given range.
*/
package main

import (
	"fmt"
	"sync"
)

var (
	result = 0
	value  = 97
)

func main() {
	calculateOne(value)
	fmt.Println("====\n")
	calculate1000(1000)
}

func Sqrt(value int, send chan<- int) {
	result = value * value
	send <- result
}

func Result(value int, get <-chan int, done chan<- bool) {
	fmt.Println("The result of", value, "squared", "is", <-get)
	done <- true
}

func calculateOne(value int) {
	cal := make(chan int)
	done := make(chan bool)
	go Sqrt(value, cal)
	go Result(value, cal, done)
	<-done
}

func calculate1000(value int) {
	var wg sync.WaitGroup
	for i := 0; i < value; i++ {
		wg.Add(1)
		go func(x int) {
			//var result int = 0
			calc := make(chan int)
			done := make(chan bool)
			go Sqrt(x, calc)
			go Result(x, calc, done)
			<-done
			wg.Done()
		}(i)
	}
	wg.Wait()
}

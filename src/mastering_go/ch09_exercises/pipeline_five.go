package main

import (
	"fmt"
	"math/rand"
	//"os"
	//"strconv"
	"time"
)

var CLOSEA = false

var DATA = make(map[int]bool)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func first(min, max int, out chan<- int) {
	for {
		if CLOSEA {
			close(out)
			return
		}
		out <- random(min, max)
	}
}

func second(out chan<- int, in <-chan int) {
	fmt.Printf("in: %s - len(%v) \n", in, len(in))
	for x := range in {
		fmt.Print(x, " ")
		_, ok := DATA[x]
		if ok {
			CLOSEA = true
		} else {
			DATA[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
}

func third(in <-chan int) {
	var sum int
	sum = 0
	for x2 := range in {
		sum = sum + x2
	}
	fmt.Printf("The sum of the random numbers is %d.\n", sum)
}

func four(send chan<- int, howMany, min, max int) {
	for i := 1; i <= howMany; i++ {
		rnd := random(min, max)
		fmt.Printf("%v: %v \n ", i, rnd)
		send <- rnd
	}
	fmt.Println()
	close(send)
}

func five(get <-chan int, send chan<- int) {
	fmt.Println("Passing a random numbers.")
	for x := range get {
		fmt.Printf("%v ", x)
		send <- x
	}
	fmt.Println()
	close(send)
}

func six(get <-chan int) {
	var sum int
	sum = 0
	for x2 := range get {
		fmt.Printf("%v ", x2)
		sum = sum + x2
	}
	fmt.Printf("\nThe sum of the random numbers is %d.\n", sum)
}

func main() {

	//if len(os.Args) != 3 {
	//	fmt.Println("Need two integer parameters!")
	//	return
	//}

	//n1, _ := strconv.Atoi(os.Args[1])
	//n2, _ := strconv.Atoi(os.Args[2])

	n1 := 28
	n2 := 100
	if n1 > n2 {
		fmt.Printf("%d should be smaller than %d.\n", n1, n2)
		return
	}

	rand.Seed(time.Now().UnixNano())
	A := make(chan int)
	B := make(chan int)

	C := make(chan int)
	D := make(chan int)

	go first(n1, n2, A)
	go second(B, A)
	third(B)

	fmt.Println("================================================================")
	go four(C, 10, n1, n2)
	go five(C, D)
	six(D)
}

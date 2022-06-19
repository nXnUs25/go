package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go hello(&wg)
	go goodbye2(&wg)
	wg.Wait()
	fmt.Println()

	wg.Add(1)
	go hello(&wg)
	wg.Wait()
	//goodbye2(&wg) we cannot defer waiting group since counter is 0 here so we need to use the other version of goodbye function
	goodbye()

	fmt.Println()
	wg.Add(1)
	go hello(&wg)
	goodbye() // this way we have this function returning output not in order we could see first goodbye message or hello
	wg.Wait()

}

func hello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello, world!")
}

func goodbye2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Good Bey...")
}

func goodbye() {
	fmt.Println("Good Bey...")
}

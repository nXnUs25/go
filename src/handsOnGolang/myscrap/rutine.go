package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Outside go ruting value is [%v] [%[1]T]\n", i)
		go func() {
			fmt.Printf("Inside go ruting value is [%v] [%[1]T]\n", i)
		}()
		time.Sleep(3 * time.Second)
	}
}

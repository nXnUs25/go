package main

import (
	"fmt"
	"time"
)

func main() {
	// create a channel
	ch := make(chan string, 1) // async channel, buffered channel ...holds values in underlaying array and values can be processed later
	ch2 := make(chan string)   // sync channel. unbuffered channel ... cannot hold any values and both sender and receiver must be present at the same time.
	// start the greeter to provide a greeting
	go greet(ch)   // first channel will send the message to sender straightaway before we reach to 1st sleep function
	go greet2(ch2) // this channel since is unbuffered will wait for the sender to send the message to reciver
	// sleep for a long time
	fmt.Println("Sleep for 5s..")
	time.Sleep(5 * time.Second)
	fmt.Println("Main ready!")
	// receive greeting
	greeting := <-ch   // it will not wait since is async
	greeting2 := <-ch2 // will wait till this line to print Greeter completed
	// sleep and print
	fmt.Println("Sleep for 1s..")
	time.Sleep(1 * time.Second)
	fmt.Println("Greeting received!")
	fmt.Println(greeting)
	fmt.Println("[2] -", greeting2)

}

// greet writes a greet to the given channel and then says goodbye
func greet(ch chan string) {
	fmt.Printf("Greeter ready!\nGreeter waiting to send greeting...\n")
	// greet
	ch <- "Hello, world!"
	fmt.Println("Greeter completed!")
}

// greet writes a greet to the given channel and then says goodbye
func greet2(ch chan<- string) {
	fmt.Printf("[2] - Greeter ready!\n[2] - Greeter waiting to send greeting...\n")
	// greet
	ch <- "[2] - Hello, world!"
	fmt.Println("[2] - Greeter completed!")
}

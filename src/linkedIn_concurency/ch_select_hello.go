package main

import (
	"fmt"
	"sync"
	"time"
)

// greetings in many languages
var hellos = []string{"[H] - Hello!", "[H] - Ciao!", "[H] - Hola!", "[H] - Hej!", "[H] - Salut!"}
var goodbyes = []string{"[G] - Goodbye!", "[G] - Arrivederci!", "[G] - Adios!", "[G] - Hej Hej!", "[G] - La revedere!"}

func main() {
	// create a channel
	ch := make(chan string, 1)
	ch2 := make(chan string, 1)
	// start the greeter to provide a greeting
	go greet(hellos, ch)
	go greet(goodbyes, ch2)
	// sleep for a long time
	time.Sleep(1 * time.Second)
	fmt.Println("Main ready!")
loop:
	for {
		select {
		case gr, ok := <-ch:
			if !ok {
				ch = nil
				break
			}
			printGreeting(gr)
		case gr2, ok := <-ch2:
			if !ok {
				ch2 = nil
				break
			}
			printGreeting(gr2)
		default:
			fmt.Println("default")
			//return
			break loop
		}
	}
	fmt.Println("\n============\n")
	ch3 := make(chan string, 1)
	ch4 := make(chan string, 1)
	go greet(hellos, ch3)
	go greet(goodbyes, ch4)
	fmt.Println("Sleep for 5s..")
	time.Sleep(5 * time.Second)
	fmt.Println("Main [2] ready!")
	fmt.Println("\n============ Hello ===============\n")
	PrintGreetings(ch3)
	fmt.Println("\n============ Goodbys ===============\n")
	PrintGreetings(ch4)
}

// greet writes a greet to the given channel and then says goodbye
func greet(greetings []string, ch chan<- string) {
	var wg sync.WaitGroup
	fmt.Println("Greeter ready!")
	// greet
	go func() {
		wg.Add(len(greetings))
		for _, g := range greetings {
			ch <- g
			wg.Done()
		}
		close(ch)
	}()
	wg.Wait()
	fmt.Println("Greeter completed!")
}

// printGreeting sleeps and prints the greeting given
func printGreeting(greeting string) {
	// sleep and print
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Greeting received!", greeting)
}

func PrintGreetings(in <-chan string) { // popline example, we passing only chan var and we can recive all the messages
done:
	for {
		select {
		case greet, ok := <-in:
			if !ok {
				in = nil
				break
			}
			printGreeting(greet)
		default:
			break done
		}
	}
}

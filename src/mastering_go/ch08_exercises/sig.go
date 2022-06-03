package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
Create a Go program that handles any three signals you choose.

(Page 448).
*/

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan) // no signal specified so all will be handled
	go func() {
		sig := <-sigChan
		handleSig(sig)
		switch sig {
		case os.Interrupt:
			handleSig(sig)
		case syscall.SIGTERM:
			handleSig(sig)
			os.Exit(0)
		case syscall.SIGUSR2:
			fmt.Println("Handling syscall.SIGUSR2!")
		case syscall.SIGALRM:
			handleSig(sig)
		case syscall.SIGHUP:
			fmt.Println("HUP catch it")
		default:
			fmt.Println("Ignoring:", sig)
		}
	}()

	for {
		fmt.Printf(".")
		time.Sleep(30 * time.Second)
	}
}

func handleSig(signal os.Signal) {
	fmt.Printf("Caught signale %v\n", signal)
}

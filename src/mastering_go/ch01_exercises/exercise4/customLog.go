package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var LOGFILE1 = "/tmp/mGo1.log"
var LOGFILE2 = "/tmp/mGo2.log"

func main() {
	f1, err := os.OpenFile(LOGFILE1, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f1.Close()

	f2, err := os.OpenFile(LOGFILE2, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f2.Close()

	/*
	   io.MultiWriter is used to combine several writers together. Here, it creates a writer w - a write to w will go to os.Stdout as well as two files
	   log.New lets us create a new log.Logger object with a custom writer
	   The log.Logger object can be passed around to functions and used by them to log things

	   [https://stackoverflow.com/questions/54820380/how-to-write-logs-to-multiple-log-files-in-golang]
	*/
	w := io.MultiWriter(os.Stdout, f1, f2)

	// LstdFlags = Ldate | Ltime
	// initial values for the standard logger
	iLog := log.New(w, "customLog ", log.LstdFlags)
	//iLog.SetOutput(w)
	iLog.Println("Hello there!")
	iLog.Println("Another log entry!")
}

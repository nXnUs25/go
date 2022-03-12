package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	buf bytes.Buffer
	l   = log.New(&buf, "KeyValue: ", log.LstdFlags|log.Lshortfile)
	I   = "[INFO] "
	D   = "[DEBUG] "
	E   = "[ERROR] "
)

type myElement struct {
	Name    string
	Surname string
	Id      string
}

var DATA = make(map[string]myElement)

func ADD(k string, n myElement) bool {
	l.Printf(I+" <ADD> Key [%s] - myElement:[%v]\n", k, n)
	if k == "" {
		return false
	}

	if LOOKUP(k) == nil {
		DATA[k] = n
		return true
	}
	return false
}

func DELETE(k string) bool {
	l.Printf(I+" <DELETE> Key [%s]\n", k)
	if LOOKUP(k) != nil {
		l.Printf(D+" <DELETE> Key [%s] - OK:[%v]\n", k, true)
		delete(DATA, k)
		return true
	}
	l.Printf(D+" <DELETE> Key [%s] - OK:[%v]\n", k, false)
	return false
}

func LOOKUP(k string) *myElement {
	l.Printf(I+" <LOOKUP> Key [%s]\n", k)
	_, ok := DATA[k]
	if ok {
		n := DATA[k]
		l.Printf(D+" <LOOKUP> Key [%s] - OK:[%v] - myElement:[%v]\n", k, ok, n)
		return &n
	} else {
		l.Printf(D+" <LOOKUP> Key [%s] - OK:[%v] - myElement:[nil]\n", k, ok)
		return nil
	}
}

func CHANGE(k string, n myElement) bool {
	l.Printf(I+" <CHANGE> Key [%s] - myElement:[%v]\n", k, n)
	DATA[k] = n
	return true
}

func PRINT() {
	l.Println(I + " <PRINT>")
	for k, d := range DATA {
		fmt.Printf("key: %s value: %v\n", k, d)
	}
}

/*
Try to improve keyValue.go by adding logging to

added few lines of logger code to demostrate this solution .. more can be added latter for
better understanding the program flow ..
*/
func main() {
	l.SetOutput(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)
	l.Println(I + "Improved version of keyValue program started succesfully")
	l.Printf(I + "Available cmds: <PRINT|ADD|DELETE|LOOKUP|CHANGE>\n")
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		tokens := strings.Fields(text)

		switch len(tokens) {
		case 0:
			continue
		case 1:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 2:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 3:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 4:
			tokens = append(tokens, "")
		}

		switch tokens[0] {
		case "PRINT":
			PRINT()
		case "STOP":
			return
		case "DELETE":
			if !DELETE(tokens[1]) {
				fmt.Println("Delete operation failed!")
			}
		case "ADD":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !ADD(tokens[1], n) {
				fmt.Println("Add operation failed!")
			}
		case "LOOKUP":
			n := LOOKUP(tokens[1])
			if n != nil {
				fmt.Printf("%v\n", *n)
			}
		case "CHANGE":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !CHANGE(tokens[1], n) {
				fmt.Println("Update operation failed!")
			}
		default:
			fmt.Println("Unknown command – please try again!")
			l.Printf(E + "Uknowe command failed!\navailable cmds: <PRINT|ADD|DELETE|LOOKUP|CHANGE>")
		}
	}
}

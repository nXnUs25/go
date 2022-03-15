package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

/*
Develop a Go program that finds all the IPv4 addresses of a log file that downloaded ZIP files.
*/
func main() {
	pwd, _ := os.Getwd()
	utils := "/src/mastering_go/ch04_exercises/regex_tasks"
	file := pwd + utils + "/logfile.log"
	f := getLogFile(file)
	zips := readLogFile(f)

	for _, line := range zips {
		pf("IP [%16s] zip requested in line [%v]\n", GetIPFormat(line), line[:len(line)-1])
	}
}

var pf = fmt.Printf
var pln = fmt.Println
var p = fmt.Print

func findZIP(input string) string {
	zip := "\\.zip"
	requestType := "(GET|POST|PUT|DELETE|HEAD|OPTIONS|CONNECT)"
	request := "/?.*"
	pattern := requestType + " " + request + zip + " "
	match := regexp.MustCompile(pattern)
	return match.FindString(input)
}

func getLogFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		pf("error opening log file %s: %v", filename, err)
		os.Exit(1)
	}
	//defer file.Close() // u cannot close the file here

	return file
}

func readLogFile(f *os.File) []string {
	defer f.Close()
	zips := []string{}
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			pf("error reading file %s", err)
			break
		}

		zip := findZIP(line)
		if zip == "" {
			continue
		} else {
			zips = append(zips, line)
		}
	}
	return zips
}

func GetIPFormat(input string) string {
	anyIP := "([0-9]{1,3}\\.){3}[0-9]{1,3}"
	matchMe := regexp.MustCompile(anyIP)
	return matchMe.FindString(input)
}

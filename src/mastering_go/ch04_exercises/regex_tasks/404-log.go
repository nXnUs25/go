package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

/*
Develop a Go program that finds the IPv4 addresses in a log file that generated a 404 HTML error message.
*/
func main() {
	pwd, _ := os.Getwd()
	utils := "/src/mastering_go/ch04_exercises/regex_tasks"
	file := pwd + utils + "/logfile.log"
	f := getLogFile(file)
	lines_404 := readLogFile(f)

	for _, line := range lines_404 {
		pf("IP [%16s] got HTTP ERR NotFound. in line [%v]\n", GetIPFormat(line), line[:len(line)-1])
	}
}

var pf = fmt.Printf
var pln = fmt.Println
var p = fmt.Print

func find404(input string) string {
	notFound := "404"
	requestType := "(GET|POST|PUT|DELETE|HEAD|OPTIONS|CONNECT)"
	request := "/+.*"
	pattern := requestType + " " + request + " " + notFound + " "
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
	errs_404 := []string{}
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			pf("error reading file %s", err)
			break
		}

		notfound_err := find404(line)
		if notfound_err == "" {
			continue
		} else {
			errs_404 = append(errs_404, line)
		}
	}
	return errs_404
}

func GetIPFormat(input string) string {
	anyIP := "([0-9]{1,3}\\.){3}[0-9]{1,3}"
	matchMe := regexp.MustCompile(anyIP)
	return matchMe.FindString(input)
}

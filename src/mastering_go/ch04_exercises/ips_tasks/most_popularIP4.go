package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"sort"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("usage: %s logFile\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	m := make(map[string]int)

	for _, fielname := range args[1:] {
		f, err := os.Open(fielname)
		if err != nil {
			fmt.Printf("error opening %s: %v\n", fielname, err)
			os.Exit(1)
		}
		defer f.Close()

		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading %s: %v\n", fielname, err)
				break
			}

			ip := findIP(line)
			trial := net.ParseIP(ip)
			if trial.To4() == nil {
				continue
			}

			calculateIPs(ip, m)
		}
	}

	// print first results not sorted
	fmt.Println(`
	This script is doing same thing as the Unix commands sort and uniq.
	You can validate this by running the following command-line:
	> go run findIP4.go logfile.log | sort -nr | uniq -c | sort -rn -k1
	`)
	print(m)
	fmt.Println()
	results := sortResults(m)
	printResults(results)

}

func findIP(input string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)
}

func calculateIPs(input string, m map[string]int) *map[string]int {
	//fmt.Printf("1. input := %s --- map : %v\n", input, m)
	if count, ok := m[input]; ok {
		count++
		m[input] = count
		//fmt.Printf("2. count := %v --- ok : %v\n", count, ok)
	} else {
		m[input] = 1
		//fmt.Printf("3. count := %v --- ok : %v\n", count, ok)
	}

	return &m
}

func print(m map[string]int) {
	fmt.Println("Most populat Ips: ")
	c := 1
	for k, v := range m {
		fmt.Printf("%-3v. [%3v] - [%16v]\n", c, v, k)
		c++
	}
}

func printResults(results []Result) {
	for _, r := range results {
		fmt.Printf("[%3v] - [%15v]\n", r.count, r.ip)
	}
}

func sortResults(m map[string]int) []Result {
	var sortedResults []Result
	for k, v := range m {
		sortedResults = append(sortedResults, Result{k, v})
	}
	sort.Slice(sortedResults, func(i, j int) bool {
		return sortedResults[i].count > sortedResults[j].count
	})
	return sortedResults
}

type Result struct {
	ip    string
	count int
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

func findIP(input string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)
}

// Try to write a Go program that prints the invalid part or parts of an IPv4 address.
// will print full bad ip not just bad octed
func findWrongIP(input string) string {
	badOctetIPv4 := "(\\.|^)([2-9][5-9][6-9]|[3-9][0-9]{2}|0[0-9]+)([^:/-])"
	//badOctetIPv4 := "(\\.|^)([2-9][5-9][6-9]|[3-9][0-9][0-9]|0[0-9][0-9]?)($|\\.)"
	ipv4Format := badOctetIPv4 // + "\\." + badOctedIPv4 + "\\." + badOctedIPv4 + "\\." + badOctedIPv4
	matchMe := regexp.MustCompile(ipv4Format)
	return matchMe.FindString(input)
}

// just get ip format even if not valid
func GetIPFormat(input string) string {
	anyIP := "([0-9]{1,3}\\.){3}[0-9]{1,3}"
	matchMe := regexp.MustCompile(anyIP)
	return matchMe.FindString(input)
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Printf("usage: %s logFile\n", filepath.Base(arguments[0]))
		os.Exit(1)
	}

	for _, filename := range arguments[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file %s\n", err)
			os.Exit(-1)
		}
		defer f.Close()

		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file %s", err)
				break
			}

			badOctet := findWrongIP(line)
			if badOctet == "" {
				continue
			}
			invalidIP := GetIPFormat(line)
			if invalidIP != "" {
				fmt.Printf("[%15v] : [%5v] : %s", invalidIP, badOctet, line)
			}
		}
	}
}

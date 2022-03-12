package main

import (
	"fmt"
	"regexp"
	"strconv"
)

/*
Develop a Go utility that uses a regular expression in order to match integers from 200 to 400.
*/

var pf = fmt.Printf

func main() {
	pattern := "\\b(200|[2-3][0-9][0-9]|400)\\b"
	numbers := []int{1, 2, 3, 4, 5, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 200,
		201, 202, 203, 204, 210, 300, 400, 401, 600, 900}
	match := regexp.MustCompile(pattern)
	for _, v := range numbers {
		found := match.FindString(strconv.Itoa(v))
		if found == "" {
			continue
		} else {
			pf("Got value [%4v] which is in the range of [200-400]\n", v)
		}
	}
}

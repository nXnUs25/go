package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 || len(flag.Args()) > 2 {
		fmt.Fprintf(os.Stderr, "usage: tabsToSpaces file <num>\n")
		return
	}
	fmt.Println(len(flag.Args()))
	file, err := os.Open(flag.Args()[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open file %s: %v\n", flag.Args()[0], err)
	}
	defer file.Close()
	spaces, err := strconv.Atoi(flag.Args()[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot phrase spaces to int: %v\n", err)
	}
	replace := spaceBuilder(spaces)
	r := bufio.NewReader(file)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot read file: %v\n", err)
			return
		}

		rx := regexp.MustCompile("\t")
		text := rx.ReplaceAllString(line, replace)
		fmt.Fprintf(os.Stdout, "%s", text)
	}
}

func spaceBuilder(spaces int) string {
	var sb strings.Builder
	for i := 0; i < spaces; i++ {
		sb.WriteString(" ")
	}
	return sb.String()
}

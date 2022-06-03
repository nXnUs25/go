package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	flag.Parse()
	fmt.Println(flag.Args()[0])
	if len(flag.Args()) != 3 {
		fmt.Fprintln(os.Stderr, "replace_word <file> <replaceFrom> <replaceTo>")
	}

	replaceFrom := flag.Args()[1]
	replaceTo := flag.Args()[2]
	file, err := os.Open(flag.Args()[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open <file> %v\n", err)
	}
	defer file.Close()

	r := bufio.NewReader(file)
	for {
		line, err := r.ReadString('\n')
		//fmt.Println(line)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot read file %v\n", err)
			return
		}

		regex := regexp.MustCompile(replaceFrom)
		text := regex.ReplaceAllString(line, replaceTo)
		fmt.Fprintf(os.Stdout, "%s", text)
	}
}

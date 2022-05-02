package main

/*
Can you create a Go version of the wc(1) utility?
Look at the manual page of wc(1) to find out about the command-line options that it supports.

(Page 448).
*/

import (
	//"flag"
	"fmt"
	//"github.com/spf13/pflag"
	//"github.com/spf13/viper"
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
)

var filename string

type wc struct {
	lines int
	words int
	char  int
	bytes int
}

var wc_syntax string = `[%s]:
 lines: %5d
 words: %5d
 chars: %5d
 bytes: %5d`

func (c *wc) readFile(file string) {
	filename = file
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
		} else {
			c.lines += 1
			c.char += len(line)
			c.bytes += len([]byte(line))
			rx := regexp.MustCompile(`[^\s]+`)
			words := rx.FindAllString(line, -1)
			c.words += len(words)
		}
	}
}

func (c *wc) getCharsCount() int {
	if c != nil {
		return c.char
	}
	return 0
}

func (c *wc) getWordsCount() int {
	if c != nil {
		return c.words
	}
	return 0
}

func (c *wc) getLinesCount() int {
	if c != nil {
		return c.lines
	}
	return 0
}

func (c *wc) getBytesCount() int {
	if c != nil {
		return c.bytes
	}
	return 0
}

func (c *wc) String() string {
	if c != nil {
		return fmt.Sprintf(wc_syntax, filename, c.lines, c.words, c.char, c.bytes)
	}
	c = &wc{lines: 0, char: 0, bytes: 0, words: 0}
	return fmt.Sprintf(wc_syntax, filename, c.lines, c.words, c.char, c.bytes)
}

func main() {

	wc := &wc{}

	wc.readFile("./test.txt")

	fmt.Println(wc)
}

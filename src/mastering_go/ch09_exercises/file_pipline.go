package main

/*
Create a pipeline that reads text files, finds the number of occurrences of a given phrase in each text file,
and calculate the total number of occurrences of the phrase in all files.
*/
import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sync"
)

type FileRW struct {
	sync.RWMutex
}

func main() {
	file := "./file.txt"
	word := "Text"
	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[Error reading file]:\n", err)
	}
	defer f.Close()
	lines := make(chan string, 10)
	result := make(chan int, 1)

	go readLines(f, lines)
	go grep(word, lines, result)
	fmt.Printf("word [%s] occurrences in the file are: %d\n", word, <-result)
}

func readLines(f *os.File, ch chan string) {
	scan := bufio.NewScanner(f)

	for scan.Scan() {
		//fmt.Println(scan.Text())
		ch <- scan.Text()
	}
	close(ch)
}

func grep(word string, lines <-chan string, results chan int) {
	current := 0
	re := regexp.MustCompile(`\b` + word + `\b`)
	for l := range lines {
		if re.MatchString(l) {
			current++
		}
	}
	results <- current
	close(results)
}

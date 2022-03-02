package ch01_exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var END = "END"

/*
Write a Go program that finds the sum of all command-line arguments that are valid numbers.
(Page 62).
*/
func CalculateSum(a []string) int64 {
	var sum int64
	sum = 0
	for x := 0; x < len(a); x++ {
		n1, err := strconv.ParseInt(a[x], 10, 64)
		if err != nil {
			f1, err := strconv.ParseFloat(a[x], 64)
			if err == nil {
				n1 = int64(f1)
			} else {
				fmt.Println("[ERROR]: Invalid number.")
				fmt.Println(err)
				os.Exit(2)
			}
		}
		sum += n1
	}
	return sum

}

func CalculateFloatSum(a []string) float64 {
	var sum float64
	for x := 1; x < len(a); x++ {
		f1, err := strconv.ParseFloat(a[x], 64)
		if err != nil {
			fmt.Println("[ERROR]: Invalid number.")
			fmt.Println(err)
			os.Exit(2)
		}
		sum += f1
	}
	return sum
}

/*
Write a Go program that finds the average value of floating-point numbers that are given as command-line arguments.
(Page 62).
*/
func CalculateAvg(a []string) (float64, float64) {
	var avg float64
	var sum float64
	var counter int
	for x := 0; x < len(a); x++ {
		counter += 1
		f1, err := strconv.ParseFloat(a[x], 64)
		if err != nil {
			fmt.Println("[ERROR]: Invalid number.")
			fmt.Println(err)
			os.Exit(2)
		}
		sum += f1
	}
	avg = sum / float64(counter)
	return sum, avg
}

/*
Write a Go program that keeps reading integers until it gets the word END as input.
(Page 62)
*/
func ScannerInput() {
	// var f *os.File
	f := os.Stdin
	defer f.Close()

	fmt.Printf("[INPUT]: [Ctrl-C or a string <%s> to stop reading] Please enter a number.\n>>> ", END)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Print(">>> ")
		txt := parseIntToString(scanner.Text())
		if stopScanner(txt, END) {
			break
		}
	}
}

func parseIntToString(s string) string {
	n, err := strconv.Atoi(s)
	if err != nil {
		return s
	}
	return strconv.Itoa(n)
}

func stopScanner(input, stop string) bool {
	if input == stop {
		return true
	} else {
		return false
	}
}

func SetEndingCondition(s string) {
	END = s
}

func GetEndingCondition() string {
	return END
}

/*
Can you modify customLog.go in order to write its log data into two log files at the same time? You might need to read Chapter 8, Telling a UNIX System What to Do, for help.
(Page 62).
	- customLog is under exercises4
	- any other helping function will be declared here

	Verification steps
	❯ cat /tmp/*.log
		customLog 2022/03/02 07:08:15 Hello there!
		customLog 2022/03/02 07:08:15 Another log entry!
		customLog 2022/03/02 07:08:15 Hello there!
		customLog 2022/03/02 07:08:15 Another log entry!
*/
